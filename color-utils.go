package colorarty

import (
	"image/color"
	"math"
)

func isBlackOrWhite(c color.Color) bool {
	r, g, b, _ := toRGBA(c)
	if r > .91 && g > .91 && b > .91 {
		return true // white
	}
	if r < .09 && g < .09 && b < .09 {
		return true // black
	}

	return false
}

func isDarkColor(c color.Color) bool {
	r, g, b, _ := toRGBA(c)
	lum := 0.2126*r + 0.7152*g + 0.0722*b
	return lum < .5
}

func toRGBA(c color.Color) (r, g, b, a float64) {
	cr, cg, cb, ca := c.RGBA()
	r = float64(cr)
	g = float64(cg)
	b = float64(cb)
	a = float64(ca)
	r /= 0x101
	g /= 0x101
	b /= 0x101
	a /= 0x101
	r /= 255.0
	g /= 255.0
	b /= 255.0
	a /= 255.0
	return
}

func isDistinct(ca, cb color.Color) bool {
	r, g, b, a := toRGBA(ca)
	r1, g1, b1, a1 := toRGBA(cb)

	const threshold = .25 // .15?

	if math.Abs(r-r1) > threshold ||
		math.Abs(g-g1) > threshold ||
		math.Abs(b-b1) > threshold ||
		math.Abs(a-a1) > threshold {
		// check for grays, prevent multiple gray colors
		if math.Abs(r-g) < .03 && math.Abs(r-b) < .03 {
			if math.Abs(r1-g1) < .03 && math.Abs(r1-b1) < .03 {
				return false
			}
		}
		return true
	}
	return false
}

func isContrasting(fore, back color.Color) bool {
	br, bg, bb, _ := toRGBA(back)
	fr, fg, fb, _ := toRGBA(fore)

	bLum := 0.2126*br + 0.7152*bg + 0.0722*bb
	fLum := 0.2126*fr + 0.7152*fg + 0.0722*fb

	contrast := 0.

	if bLum > fLum {
		contrast = (bLum + 0.05) / (fLum + 0.05)
	} else {
		contrast = (fLum + 0.05) / (bLum + 0.05)
	}

	//return contrast > 3.0; //3-4.5 W3C recommends 3:1 ratio, but that filters too many colors
	return contrast > 1.6
}
