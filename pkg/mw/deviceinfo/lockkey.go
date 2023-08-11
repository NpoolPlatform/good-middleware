package deviceinfo

import (
	"encoding/hex"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) lockKey() string {
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateDeviceInfo, hex.EncodeToString([]byte(*h.Type)))
}
