package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

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
			img.Set(px, py, mandelbrot64(z))
		}
	}
	f, _ := os.Create("output.png")
	png.Encode(f, img)
}

var palettes = [...]color.RGBA{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[n%16]
			}
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			switch {
			case n > 50:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[n%16]
			}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 20
	const contrast = 15
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 5:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[i%16]
			}
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	const iterations = 20
	const contrast = 15
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			switch {
			case i > 5:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[i%16]
			}
		}
	}
	return color.Black
}


//go run main.go -x -2 -zoom 15 -type complex64  > output/complex64.png
//go run main.go -x -2 -zoom 15 -type complex128  > output/complex128.png
//go run main.go -x -2 -zoom 15 -type bigfloat  > output/bigfloat.png
//go run main.go -x -2 -zoom 15 -type bigrat  > output/bigrat.png
