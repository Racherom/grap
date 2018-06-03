package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

const (
	winWidth = 255
	winHight = 255
)

type (
	pixels []byte
)

func (p *pixels) set(x, y int, c color.Color) {
	i := (y*winWidth + x) * 4
	if len(*p) > i+3 {
		c := color.RGBAModel.Convert(c)
		rgba := c.(color.RGBA)
		(*p)[i] = byte(rgba.B)
		(*p)[i+1] = byte(rgba.G)
		(*p)[i+2] = byte(rgba.R)
		(*p)[i+3] = byte(rgba.A)
	}
}

func main() {

	win, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, winWidth, winHight)
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()

	p := make(pixels, winHight*winWidth*4)
	for y := 0; y < winHight; y++ {
		for x := 0; x < winWidth; x++ {
			p.set(x, y, color.RGBA{0, uint8(x % 255), uint8(y % 255), 0})
		}
	}
	tex.Update(nil, p, int(winWidth*int32(4)))
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	sdl.Delay(5000)
}
