package topmost

import (
	"fmt"
	"time"
)

func (h *Handler) formalizeStartEnd() error {
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
