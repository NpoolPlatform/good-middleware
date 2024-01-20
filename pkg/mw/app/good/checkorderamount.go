package appgood

import (
	"fmt"
)

func (h *Handler) checkOrderAmount() error {
	if h.MinOrderAmount == nil || h.MaxOrderAmount == nil {
		return nil
	}
	if h.MinOrderAmount.Cmp(*h.MaxOrderAmount) > 0 {
		return fmt.Errorf("invalid orderamount")
	}
	return nil
}
