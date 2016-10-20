package random

import "math"

func Beta(x, y float64) float64 {
	return (math.Gamma(x) * math.Gamma(y)) / math.Gamma(x+y)
}
