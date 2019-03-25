package gradient

import (
	"fmt"
	"image/color"
)

func ExampleLinearPoint() {
	a := color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	b := color.RGBA{0x00, 0x00, 0xFF, 0xFF}

	c := LinearPoint(a, b, 0.5)
	fmt.Printf("#%02x%02x%02x @ %02x", c.R, c.G, c.B, c.A)
	// Output: #80007f @ ff
}
