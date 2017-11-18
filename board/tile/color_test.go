package tile

import (
	"testing"

	"github.com/fatih/color"
)

func TestColorForNumber(t *testing.T) {
	numberToColorExpected := map[int64]color.Attribute{
		2:                color.FgRed,
		4:                color.FgGreen,
		8:                color.FgYellow,
		16:               color.FgBlue,
		32:               color.FgMagenta,
		64:               color.FgCyan,
		128:              color.FgWhite,
		256:              color.FgHiRed,
		512:              color.FgHiGreen,
		1024:             color.FgHiYellow,
		2048:             color.FgHiBlue,
		4096:             color.FgHiMagenta,
		8192:             color.FgHiCyan,
		16384:            color.FgHiWhite,
		32768:            color.FgRed,
		65536:            color.FgGreen,
		131072:           color.FgYellow,
		1125899906842624: color.FgCyan,
	}

	for n, expected := range numberToColorExpected {
		actual := ColorForNumber(n)
		if actual != expected {
			t.Errorf("ColorForNumber(%d): expected %v, actual %v", n, expected, actual)
		}
	}
}
