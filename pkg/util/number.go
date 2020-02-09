package util

import (
	"math"
	"math/rand"
)

func GenerateNumberByDigitLen(nDigit int) int {
	limit := int(math.Pow(10.0, float64(nDigit)))
	base := limit / 10
	return base + rand.Intn(limit-base-1)
}
