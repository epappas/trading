package shapeshift

import (
	"strconv"
)

// ToFloat ShapeShift's API responds in float and string for decimals for different functions.
// Since we arn't really using 'big numbers' I think it's ok to be using this.
// This golang package is not doing any math, just responding back from ShapeShift API.
func ToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}
