package tile

import (
	"math"

	"github.com/fatih/color"
)

func ColorForNumber(n int64) color.Attribute {
	exp := int(math.Log2(float64(n)))
	exp = exp % 15

	if n >= 32768 {
		exp += 1
	}

	if exp < 8 {
		return color.Attribute(exp + int(color.FgBlack))
	} else {
		return color.Attribute(exp - 8 + int(color.FgHiBlack) + 1)
	}
}
