package colorarty

import (
	"image"
	"image/color"
	"sort"
)

// Analyse performs several operations on a given image in order to extract
// background, primary, secondary and detail colors for it.
func Analyse(img image.Image) *Result {
	colors := make(colorCounter)
	res := Result{}
	res.BackgroundColor = findEdgeColor(img, &colors)

	res.PrimaryColor, res.SecondaryColor, res.DetailColor = findTextColors(colors, res.BackgroundColor)
	hasDarkBackground := isDarkColor(*res.BackgroundColor)
	replacementColor := color.Color(color.RGBA{0, 0, 0, 255})
	if hasDarkBackground {
		replacementColor = color.Color(color.RGBA{255, 255, 255, 255})
	}

	if res.PrimaryColor == nil {
		res.PrimaryColor = &replacementColor
	}
	if res.SecondaryColor == nil {
		res.SecondaryColor = &replacementColor
	}
	if res.DetailColor == nil {
		res.DetailColor = &replacementColor
	}
	return &res
}

func addItem(i *color.Color, to *colorCounter) {
	toValue := *to
	colorValue := *i

	r, g, b, a := colorValue.RGBA()

	key := r + g + b + a

	if v, ok := toValue[key]; ok {
		v.Count++
		toValue[key] = v
	} else {
		toValue[key] = countedColor{
			Color: colorValue,
			Count: 1,
		}
	}
	to = &toValue
}

func findEdgeColor(img image.Image, colors *colorCounter) *color.Color {
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()

	leftEdgeColors := make(colorCounter)
	searchColumnX := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := img.At(x, y)
			_, _, _, ca := toRGBA(c)
			if x == searchColumnX {
				// make sure it's a meaningful colour
				if ca > .5 {
					addItem(&c, &leftEdgeColors)
				}
			}
			if ca > 0 {
				addItem(&c, colors)
			}
		}
		// background is clear, keep looking in next column for background color
		if len(leftEdgeColors) == 0 {
			searchColumnX++
		}
	}

	var sortedColors countedColors

	for _, data := range leftEdgeColors {
		randomColorsThreshold := uint(float64(height) * 0.01)

		if data.Count <= randomColorsThreshold { // prevent using random colors, threshold based on input image height
			continue
		}

		sortedColors = append(sortedColors, data)
	}

	sort.Sort(sortedColors)

	var proposedEdgeColor *countedColor
	if len(sortedColors) > 0 {
		proposedEdgeColor = &sortedColors[0]
		if isBlackOrWhite(proposedEdgeColor.Color) { // want to choose color over black/white so we keep looking
			for i := 1; i < len(sortedColors); i++ {
				nextProposedColor := sortedColors[i]
				if float64(nextProposedColor.Count/proposedEdgeColor.Count) > .3 { // make sure the second choice color is 30% as common as the first choice
					if !isBlackOrWhite(nextProposedColor.Color) {
						proposedEdgeColor = &nextProposedColor
						break
					}
				} else {
					// reached color threshold less than 40% of the original proposed edge color so bail
					break
				}
			}
		}
	}
	if proposedEdgeColor != nil {
		return &proposedEdgeColor.Color
	}
	return nil
}

func findTextColors(counter colorCounter, backgroundColor *color.Color) (primary, secondary, detail *color.Color) {
	if backgroundColor == nil {
		return nil, nil, nil // whoa.
	}
	findDarkTextColor := !isDarkColor(*backgroundColor)
	var sortedColors countedColors

	for _, data := range counter {
		if isDarkColor(data.Color) == findDarkTextColor {
			sortedColors = append(sortedColors, data)
		}
	}

	sort.Sort(sortedColors)

	for _, c := range sortedColors {
		func(c countedColor) {
			if primary == nil {
				if isContrasting(c.Color, *backgroundColor) {
					primary = &c.Color
				}
			} else if secondary == nil {
				if isDistinct(c.Color, *primary) && isContrasting(c.Color, *backgroundColor) {
					secondary = &c.Color
				}
			} else if detail == nil {
				if isDistinct(c.Color, *secondary) && isDistinct(c.Color, *primary) && isContrasting(c.Color, *backgroundColor) && detail == nil {
					detail = &c.Color
				}
			}
		}(c)
	}
	return
}
