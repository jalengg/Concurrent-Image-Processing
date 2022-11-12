// Package png allows for loading png images and applying
// image flitering effects on them.
package png

import "image/color"

// Grayscale applies a grayscale filtering effect to the image
func (img *Image) Grayscale() {

	// Bounds returns defines the dimensions of the image. Always
	// use the bounds Min and Max fields to get out the width
	// and height for the image
	bounds := img.Out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//Returns the pixel (i.e., RGBA) value at a (x,y) position
			// Note: These get returned as int32 so based on the math you'll
			// be performing you'll need to do a conversion to float64(..)
			r, g, b, a := img.In.At(x, y).RGBA()

			//Note: The values for r,g,b,a for this assignment will range between [0, 65535].
			//For certain computations (i.e., convolution) the values might fall outside this
			// range so you need to clamp them between those values.
			greyC := clamp(float64(r+g+b) / 3)

			//Note: The values need to be stored back as uint16 (I know weird..but there's valid reasons
			// for this that I won't get into right now).
			img.Out.Set(x, y, color.RGBA64{greyC, greyC, greyC, uint16(a)})
		}
	}
}

// Sharpen applies a sharpen filtering effect to the image
func (img *Image) Sharpen() {
	s := [9]float64{0,-1,0,-1,5,-1,0,-1,0}
	bounds := img.Out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.convolution(x, y, s)
		}
	}
}

// Edge applies a edge detection filtering effect to the image
func (img *Image) Edge() {
	e := [9]float64{-1,-1,-1,-1,8,-1,-1,-1,-1}
	bounds := img.Out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.convolution(x, y, e)
		}
	}
}

// Blur performs a blur effect with the following kernel 
func (img *Image) Blur() {
	b := [9]float64{1 / 9.0, 1 / 9, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0}
	bounds := img.Out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.convolution(x, y, b)
		}
	}
}

func (img *Image) convolution(x int, y int, matrix [9]float64) {
	var r, g, b, a [9]uint32
	var rNew, gNew, bNew [9]uint16
	bounds := img.Out.Bounds()
	xMin := bounds.Min.X
	yMin := bounds.Min.Y

	for i := -1; i < 1; i++ {
		for j := -1; j < 1; j++ {
			pos := (i+2)*(j+2)-1
			if x + i < xMin || y + j < yMin {
				r[pos] = 0
				g[pos] = 0
				b[pos] = 0
				a[pos] = 0
			} else {
				r[pos], g[pos], b[pos], a[pos] = img.In.At(x+i, y+j).RGBA()
			}
			rNew[pos] = clamp(float64(r[pos]) * matrix[pos])
			gNew[pos] = clamp(float64(g[pos]) * matrix[pos])
			bNew[pos] = clamp(float64(b[pos]) * matrix[pos])
			img.Out.Set(x, y, color.RGBA64{rNew[pos], gNew[pos], bNew[pos], uint16(a[pos])})
		}
	}
}
