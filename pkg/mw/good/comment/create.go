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
	appgood *ent.AppGood
}

func (h *createHandler) getAppGood(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		appGood, err := cli.AppGood.Get(ctx, *h.AppGoodID)
		if err != nil {
			return err
		}
		if appGood == nil {
			return fmt.Errorf("app good not found %v", *h.AppGoodID)
		}
		h.appgood = appGood
		return nil
	})
}

func (h *createHandler) createComment(ctx context.Context, tx *ent.Tx) error {
	if h.ReplyToID != nil {
		comment, err := tx.Comment.Get(ctx, *h.ReplyToID)
		if err != nil {
			return err
		}
		if comment == nil {
			return fmt.Errorf("comment not found")
		}
		if comment.AppGoodID != *h.AppGoodID {
			return fmt.Errorf("appgoodid not matched")
		}
	}
	if _, err := commentcrud.CreateSet(
		tx.Comment.Create(),
		&commentcrud.Req{
			ID:        h.ID,
			AppID:     h.AppID,
			UserID:    h.UserID,
			GoodID:    &h.appgood.GoodID,
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
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: h.appgood.GoodID},
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
	if err := handler.getAppGood(ctx); err != nil {
		return nil, err
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
