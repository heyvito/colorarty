# colorarty
[![Go Report Card](https://goreportcard.com/badge/github.com/victorgama/colorarty)](https://goreportcard.com/report/github.com/victorgama/colorarty)
[![GoDoc](https://godoc.org/github.com/victorgama/go-unfurl?status.svg)](https://godoc.org/github.com/victorgama/go-unfurl)

**colorarty** is a small library that analyses images and extracts a background, primary, secondary and detail colors, all suitable for reading.

Here's a simple example that converts extracted colors into CSS declarations:
```go
package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"

	"github.com/victorgama/colorarty"
)

func main() {
	reader, _ := os.Open("/Users/victorgama/Downloads/sample.jpeg")
	img, _, _ := image.Decode(reader)
	result := colorarty.Analyse(img)
	fmt.Printf("#background { background: %s }\n", toCSS(*result.BackgroundColor))
	fmt.Printf(".primary { color: %s }\n", toCSS(*result.PrimaryColor))
	fmt.Printf(".secondary { color: %s }\n", toCSS(*result.SecondaryColor))
	fmt.Printf(".detail { color: %s }\n", toCSS(*result.DetailColor))
}

func toCSS(c color.Color) string {
	cr, cg, cb, _ := c.RGBA()
	r := float64(cr)
	g := float64(cg)
	b := float64(cb)
	r /= 0x101
	g /= 0x101
	b /= 0x101
	return fmt.Sprintf("rgba(%.0f, %.0f, %.0f, 1)", r, g, b)
}

```

The generated CSS was used to build this next example:

![example](https://www.dropbox.com/s/he76t7l20214lf3/colorarty-demo.png?dl=1)

## Installing
1. Download and install it:
```
$ go get -u github.com/victorgama/colorarty
```
2. Import it in your code:
```
import "github.com/victorgama/colorarty"
```

## License

**colorarty** was inspired by [this](https://panic.com/blog/itunes-11-and-colors/) post and [this](https://github.com/panicinc/ColorArt) project, both from Panic, Inc. 💖🦄

```
MIT License

Copyright (c) 2016 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
