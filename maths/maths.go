package maths

import (
	"math"
	"math/big"
)

// Round float with the given precision
// with number = 33.99 and precision = 1, Round returns 34
// with number = 33.99 and precision = 2, Round returns 33.99
// with number = 33.999 and precision = 2, Round returns 33.99
// with number = 33.99 and precision = 3, Round returns 33.999
func Round(number float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(number*ratio) / ratio
}

// Truncate float with the given precision
// with number = 33.999 and precision = 0.1, Round returns 33.9
// with number = 33.999 and precision = 0.01, Round returns 33.99
// with number = 33.999 and precision = 0.001, Round returns 33.999
func Truncate(number float64, precision float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(number)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(precision)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	number, _ = bf.Mul(bf, bu).Float64()
	return number
}
