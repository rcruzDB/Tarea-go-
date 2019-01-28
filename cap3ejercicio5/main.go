package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xf4, 0x43, 0x36, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
	color.RGBA{0x4c, 0xaf, 0x50, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
	color.RGBA{0x21, 0x96, 0xf3, 0xff}, 
	color.RGBA{0xff, 0xff, 0xff, 0xff}, 
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) 
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%len(palette)]
		}
	}
	return color.Black
}