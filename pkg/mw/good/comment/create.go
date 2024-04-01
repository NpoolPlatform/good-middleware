package comment

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/comment"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	appgood           *appgoodpb.Good
	orderFirstComment bool
}

func (h *createHandler) getAppGood(ctx context.Context) error {
	handler, err := appgood1.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.EntID = h.AppGoodID
	info, err := handler.GetGood(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}
	h.appgood = info
	return nil
}

func (h *createHandler) createComment(ctx context.Context, tx *ent.Tx) error {
	if h.ReplyToID != nil {
		handler, err := NewHandler(ctx)
		if err != nil {
			return err
		}
		handler.EntID = h.ReplyToID
		comment, err := handler.GetComment(ctx)
		if err != nil {
			return err
		}
		if comment == nil {
			return fmt.Errorf("comment not found")
		}
		if comment.AppGoodID != h.AppGoodID.String() {
			return fmt.Errorf("appgoodid not matched")
		}
	}

	goodid := uuid.MustParse(h.appgood.GoodID)
	req := &commentcrud.Req{
		EntID:             h.EntID,
		AppID:             h.AppID,
		UserID:            h.UserID,
		GoodID:            &goodid,
		AppGoodID:         h.AppGoodID,
		OrderID:           h.OrderID,
		Content:           h.Content,
		ReplyToID:         h.ReplyToID,
		Anonymous:         h.Anonymous,
		PurchasedUser:     h.PurchasedUser,
		TrialUser:         h.TrialUser,
		OrderFirstComment: &h.orderFirstComment,
	}
	if h.orderFirstComment {
		req.Score = h.Score
	}

	if _, err := commentcrud.CreateSet(tx.Comment.Create(), req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateGoodComment(ctx context.Context, tx *ent.Tx) error {
	goodid := uuid.MustParse(h.appgood.GoodID)
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: goodid},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	commentCount := info.CommentCount + 1
	req := &extrainfocrud.Req{
		CommentCount: &commentCount,
	}
	if h.orderFirstComment && h.Score != nil {
		scoreCount := info.ScoreCount + 1
		req.ScoreCount = &scoreCount
		score := h.Score.Add(
			info.Score.Mul(
				decimal.NewFromInt(int64(info.ScoreCount)),
			),
		).Div(decimal.NewFromInt(int64(scoreCount)))
		req.Score = &score
	}

	if _, err := extrainfocrud.UpdateSet(info.Update(), req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateComment(ctx context.Context) (*npool.Comment, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCommentGood, *h.AppID, *h.UserID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getAppGood(ctx); err != nil {
		return nil, err
	}

	if h.OrderID != nil {
		handler.Conds = &commentcrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
			OrderID:   &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID},
		}
		exist, err := handler.ExistCommentConds(ctx)
		if err != nil {
			return nil, err
		}
		handler.orderFirstComment = !exist
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createComment(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateGoodComment(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
