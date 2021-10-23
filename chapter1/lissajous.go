package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	color.RGBA{R: 0x70, G: 0xf2, B: 0x38, A: 0xff},
	color.RGBA{R: 0x7c, G: 0x00, B: 0x5f, A: 0xff},
	color.RGBA{R: 0x05, G: 0x40, B: 0xf4, A: 0xff},
	color.RGBA{R: 0xde, G: 0x54, B: 0x18, A: 0xff},
	color.RGBA{R: 0xff, G: 0xf2, B: 0x00, A: 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)


//func main() {
//	Lissajous(os.Stdout)
//}


func Lissajous(out io.Writer, cycles float64) {
	const (
		//cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(4)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
