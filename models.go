package colorarty

import "image/color"

// Result contains the extracted picture colors for the input image
type Result struct {
	BackgroundColor *color.Color
	PrimaryColor    *color.Color
	SecondaryColor  *color.Color
	DetailColor     *color.Color
}

type countedColor struct {
	Color color.Color
	Count uint
}

type colorCounter map[uint32]countedColor

type countedColors []countedColor

func (c countedColors) Len() int           { return len(c) }
func (c countedColors) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c countedColors) Less(i, j int) bool { return c[i].Count > c[j].Count }
