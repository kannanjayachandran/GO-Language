package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // Black
}

const (
	whiteIdx = 0
	blackIdx = 1
)

func main() {

	f, err := os.Create("Out.gif")
	if err != nil {
		panic("File not created")
	}
	defer f.Close()
	drawShape(f)

}

func drawShape(out *os.File) {

	const (
		cycles  = 6     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nFrames = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0

	for i := 0; i < nFrames; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {

			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)

}
