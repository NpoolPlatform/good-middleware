package good

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) lockKey() string {
	sha := sha256.Sum224([]byte(*h.Title))
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateGood, hex.EncodeToString(sha[:]))
}
