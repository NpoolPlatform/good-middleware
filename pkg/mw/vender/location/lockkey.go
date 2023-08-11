package location

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) lockKey() string {
	str := fmt.Sprintf("%v:%v:%v:%v", *h.Country, *h.Province, *h.City, *h.Address)
	sha := sha256.Sum224([]byte(str))
	return fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateVendorLocation,
		hex.EncodeToString(sha[:]),
		*h.BrandID,
	)
}
