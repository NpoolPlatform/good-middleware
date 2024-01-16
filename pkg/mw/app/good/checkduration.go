package appgood

import (
	"fmt"
)

func (h *Handler) checkDuration() error {
	if h.MinOrderDuration == nil || h.MaxOrderDuration == nil {
		return nil
	}
	if *h.MinOrderDuration > *h.MaxOrderDuration {
		return fmt.Errorf("invalid duration")
	}
	return nil
}
