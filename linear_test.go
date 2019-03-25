package gradient

import (
	"image/color"
	"testing"
)

var testCases = []struct {
	min uint8
	max uint8
	pos float64
	val uint8
}{
	{min: 0x00, max: 0xFF, pos: 0, val: 0x00},
	{min: 0x00, max: 0xFF, pos: 1, val: 0xFF},
	{min: 0xFF, max: 0x00, pos: 0, val: 0xFF},
	{min: 0xFF, max: 0x00, pos: 1, val: 0x00},
	{min: 0xFF, max: 0x00, pos: 2, val: 0x00},
	{min: 0xFF, max: 0x00, pos: -1, val: 0xFF},
	{min: 0x00, max: 0xFF, pos: 0.5, val: 0x7F},
	{min: 0xFF, max: 0x00, pos: 0.5, val: 0x80},
}

func TestLinear(t *testing.T) {

	t.Run("red channel", func(st *testing.T) {
		for _, tc := range testCases {
			a := color.RGBA{tc.min, 0x00, 0x00, 0x00}
			b := color.RGBA{tc.max, 0x00, 0x00, 0x00}
			c := LinearPoint(a, b, tc.pos)

			if got, want := c.R, tc.val; got != want {
				st.Errorf("min %02x, max: %02x, pos: %2f, got: %02x, want: %02x", tc.min, tc.max, tc.pos, got, want)
			}
		}
	})

	t.Run("green channel", func(st *testing.T) {
		for _, tc := range testCases {
			a := color.RGBA{0x00, tc.min, 0x00, 0x00}
			b := color.RGBA{0x00, tc.max, 0x00, 0x00}
			c := LinearPoint(a, b, tc.pos)

			if got, want := c.G, tc.val; got != want {
				st.Errorf("min %02x, max: %02x, pos: %2f, got: %02x, want: %02x", tc.min, tc.max, tc.pos, got, want)
			}
		}
	})

	t.Run("blue channel", func(st *testing.T) {
		for _, tc := range testCases {
			a := color.RGBA{0x00, 0x00, tc.min, 0x00}
			b := color.RGBA{0x00, 0x00, tc.max, 0x00}
			c := LinearPoint(a, b, tc.pos)

			if got, want := c.B, tc.val; got != want {
				st.Errorf("min %02x, max: %02x, pos: %2f, got: %02x, want: %02x", tc.min, tc.max, tc.pos, got, want)
			}
		}
	})

	t.Run("alpha channel", func(st *testing.T) {
		for _, tc := range testCases {
			a := color.RGBA{0x00, 0x00, 0x00, tc.min}
			b := color.RGBA{0x00, 0x00, 0x00, tc.max}
			c := LinearPoint(a, b, tc.pos)

			if got, want := c.A, tc.val; got != want {
				st.Errorf("min %02x, max: %02x, pos: %2f, got: %02x, want: %02x", tc.min, tc.max, tc.pos, got, want)
			}
		}
	})
}
