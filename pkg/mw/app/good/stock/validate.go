package appstock

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func (h *Handler) validate() error {
	notNil := 0
	if h.Locked != nil {
		if positive && h.Locked.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid locked")
		}
		notNil += 1
	}
	if h.WaitStart != nil {
		if positive && h.WaitStart.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid waitstart")
		}
		notNil += 1
	}
	if h.InService != nil {
		if positive && h.InService.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid inservice")
		}
		notNil += 1
	}
	if h.Reserved != nil {
		if positive && h.Reserved.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid reserved")
		}
		notNil += 1
	}
	if notNil > 1 {
		return fmt.Errorf("invalid input")
	}
	return nil
}
