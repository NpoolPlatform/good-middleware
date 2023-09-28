package comment

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/comment"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createComment(ctx context.Context, tx *ent.Tx) error {
	if _, err := commentcrud.CreateSet(
		tx.Comment.Create(),
		&commentcrud.Req{
			ID:        h.ID,
			AppID:     h.AppID,
			UserID:    h.UserID,
			AppGoodID: h.AppGoodID,
			OrderID:   h.OrderID,
			Content:   h.Content,
			ReplyToID: h.ReplyToID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateGoodComment(ctx context.Context, tx *ent.Tx) error {
	appGood, err := tx.AppGood.Get(ctx, *h.AppGoodID)
	if err != nil {
		return err
	}
	if appGood == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: appGood.GoodID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	commentCount := info.CommentCount + 1
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			CommentCount: &commentCount,
		},
	).Save(ctx); err != nil {
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
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
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
