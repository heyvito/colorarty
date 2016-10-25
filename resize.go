package colorarty

import (
	"image"

	re "github.com/nfnt/resize"
)

func resize(img image.Image, sq image.Point) image.Image {
	imgSize := img.Bounds().Size()
	var targetSize image.Point
	if imgSize.Y > imgSize.X {
		targetSize = image.Point{
			X: imgSize.X,
			Y: imgSize.X,
		}
	} else {
		targetSize = image.Point{
			X: imgSize.Y,
			Y: imgSize.Y,
		}
	}

	return re.Resize(uint(targetSize.X), uint(targetSize.Y), img, re.Bilinear)
}
