package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

type fn func(complex128) complex128

var colorPool = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height)) // create a 1024 * 1024 canvas
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, z4(z)) // loop every pixels set a specific color
		}
	}
	png.Encode(os.Stdout, img)
}

// NOTE: 牛顿迭代法https://blog.csdn.net/weixin_42943114/article/details/121905957
func z4(z complex128) color.Color {
	// f(z) = z^4 -1
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	// f(z)/f'(z) = (z^4 -1)/(4z^3)
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func newton(z complex128, f fn, fPrime fn) color.Color {
	const iterations = 37
	// const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		//  牛顿迭代法：f(z+1) = f(z) - f(z)/f'(z)
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(colorPool) == 0 {
					panic("no colors left")
				}
				c = colorPool[0]
				colorPool = colorPool[1:]
				chosenColors[root] = c
			}
			// NOTE: 下面如果处理图像会稍有变化，，见out_1.png
			// NOTE: color.RGBToYCbCr将RGB三元组转换为Y’CbCr三元组
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			// NOTE: 求y= y - y*ln(i)/ln(iterations) 亮度调整，每个起点到四个根的迭代次数对应阴影的灰度。
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

// TODO: round函数 添加一个小的偏移量，然后保留digits位小数？
func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	// NOTE: math.Trunc取整 math.Copysign返回值±0.5，符号取决于f
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
