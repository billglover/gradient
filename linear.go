package gradient

import "image/color"

// LinearPoint takes two colors and a float64, `p`, representing a position
// between them. It returns the colour represented by position `p` assuming a
// linear interpolation between the two colors. The position `p` is capped to
// the range `0 <= p <= 1`.
func LinearPoint(a, b color.Color, p float64) color.RGBA {
	if p < 0 {
		p = 0
	}
	if p > 1 {
		p = 1
	}

	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()

	red := uint8(int(float64(int(r2>>8)-int(r1>>8))*p) + int(r1>>8))
	green := uint8(int(float64(int(g2>>8)-int(g1>>8))*p) + int(g1>>8))
	blue := uint8(int(float64(int(b2>>8)-int(b1>>8))*p) + int(b1>>8))
	alpha := uint8(int(float64(int(a2>>8)-int(a1>>8))*p) + int(a1>>8))

	c := color.RGBA{red, green, blue, alpha}
	return c
}
