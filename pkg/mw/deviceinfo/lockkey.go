package deviceinfo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) lockKey() string {
	sha := sha256.Sum224([]byte(*h.Type))
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateDeviceInfo, hex.EncodeToString(sha[:]))
}
