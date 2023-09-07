package topmost

import (
	"context"
	"fmt"
	"time"
)

func (h *Handler) formalizeStartEnd(ctx context.Context) error {
	if h.ID != nil {
		info, err := h.GetTopMost(ctx)
		if err != nil {
			return err
		}
		if h.EndAt == nil {
			h.EndAt = &info.EndAt
		}
		if h.StartAt == nil {
			h.StartAt = &info.StartAt
		}
	}
	if h.EndAt == nil {
		return fmt.Errorf("invalid endat")
	}
	now := uint32(time.Now().Unix())
	if h.StartAt == nil {
		h.StartAt = &now
	}
	if *h.EndAt <= *h.StartAt {
		return fmt.Errorf("invalid startend")
	}
	return nil
}
