package maths

import (
	"math"
)

// Round float with the given precision
// with number = 33.99 and precision = 1, Round returns 34
// with number = 33.99 and precision = 2, Round returns 33.99
// with number = 33.999 and precision = 2, Round returns 33.99
// with number = 33.99 and precision = 3, Round returns 33.999
func round(k float64) int {
  return int(k + math.Copysign(0.001, k))
}

// truncate return the given float64 value with a precision decimal
// It does not round it to .1, example:
// with num = 1.123 & precision = 1, Truncate returns: 1.1
// with num = 33.99 & precision = 2, Truncate returns: 33.99
// with num = 33.99  & precision = 1, Truncate returns: 33.9
func truncate(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

