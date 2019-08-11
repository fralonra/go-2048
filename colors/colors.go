// credit to https://github.com/hajimehoshi/ebiten/blob/master/examples/2048/

package colors

import (
	"image/color"
)

// Get the cell color using default preset.
func CellColor(value int) color.Color {
	switch value {
	case 0:
		return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	case 2:
		return color.RGBA{0xee, 0xe4, 0xda, 0xff}
	case 4:
		return color.RGBA{0xed, 0xe0, 0xc8, 0xff}
	case 8:
		return color.RGBA{0xf2, 0xb1, 0x79, 0xff}
	case 16:
		return color.RGBA{0xf5, 0x95, 0x63, 0xff}
	case 32:
		return color.RGBA{0xf6, 0x7c, 0x5f, 0xff}
	case 64:
		return color.RGBA{0xf6, 0x5e, 0x3b, 0xff}
	case 128:
		return color.RGBA{0xed, 0xcf, 0x72, 0xff}
	case 256:
		return color.RGBA{0xed, 0xcc, 0x61, 0xff}
	case 512:
		return color.RGBA{0xed, 0xc8, 0x50, 0xff}
	case 1024:
		return color.RGBA{0xed, 0xc5, 0x3f, 0xff}
	case 2048:
		return color.RGBA{0xed, 0xc2, 0x2e, 0xff}
	}
	return color.White
}

// Transform colors to separated r, g, b, a values.
func ColorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	rf := float64(r) / 0xffff
	gf := float64(g) / 0xffff
	bf := float64(b) / 0xffff
	af := float64(a) / 0xffff
	// Convert to non-premultiplied alpha components.
	if 0 < af {
		rf /= af
		gf /= af
		bf /= af
	}
	return rf, gf, bf, af
}
