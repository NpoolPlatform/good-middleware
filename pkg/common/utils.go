package common

import (
	"math"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

func Seconds2Durations(seconds uint32, durationType types.GoodDurationType) uint32 {
	switch durationType {
	case types.GoodDurationType_GoodDurationByHour:
		return uint32(math.Ceil(float64(seconds) / float64(timedef.SecondsPerHour)))
	case types.GoodDurationType_GoodDurationByDay:
		return uint32(math.Ceil(float64(seconds) / float64(timedef.SecondsPerDay)))
	case types.GoodDurationType_GoodDurationByMonth:
		return uint32(math.Ceil(float64(seconds) / float64(timedef.SecondsPerMonth)))
	case types.GoodDurationType_GoodDurationByYear:
		return uint32(math.Ceil(float64(seconds) / float64(timedef.SecondsPerYear)))
	}
	return seconds
}
