package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"image/color"
	"time"
)

const (
	winWidth = 800
	winHight = 600
)

type (
	pixels []byte
	pos    struct {
		x, y, xv, yv float32
	}

	entity struct {
		pos
		color color.Color
	}

	ball struct {
		entity
		radius float32
	}

	paddle struct {
		entity
		w, h int
	}
)

func (pixels) ColorModel() color.Model {
	return color.RGBAModel
}

func (p pixels) Bounds() image.Rectangle {
	return image.Rect(0, 0, winWidth-1, winHight-1)
}

func (p *pixels) At(x, y int) color.Color {
	i := (y*winWidth + x) * 4
	if len(*p) > i+3 && i >= 0 {
		return color.RGBA{
			B: (*p)[i],
			G: (*p)[i+1],
			R: (*p)[i+2],
			A: (*p)[i+3],
		}
	}
	return color.Black
}

func (p pixels) clear() {
	for i := range p {
		p[i] = 0
	}
}

func (p *paddle) draw(pix pixels) {
	startX := int(p.x) - p.w/2
	startY := int(p.y) - p.h/2
	if startY < 0 {
		startY = 0
	}
	if startY+p.h > winHight {
		startY = winHight - p.h
	}
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			pix.Set(startX+x, startY+y, p.color)
		}
	}
}
func (p *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] > 0 {
		p.y--
	}

	if keyState[sdl.SCANCODE_DOWN] > 0 {
		p.y++
	}
}

func (p *paddle) aiUpdate(b *ball) {
	p.y = b.y
}

func (b *ball) draw(pix pixels) {
	for y := -b.radius; y < b.radius; y++ {
		for x := -b.radius; x < b.radius; x++ {
			if x*x+y*y < b.radius*b.radius {
				pix.Set(int(b.x+x), int(b.y+y), b.color)
			}
		}
	}
}

func (b *ball) update(lP, rP *paddle) {
	b.x += b.xv
	b.y += b.yv

	b.color = color.RGBA{R: 0, G: uint8(int(b.y) % 255), B: uint8(int(b.x) % 255), A: 0}

	if b.y-b.radius < 0 || b.y+b.radius >= winHight {
		b.yv = -b.yv
	}

}

func (p *pixels) Set(x, y int, c color.Color) {
	i := (y*winWidth + x) * 4
	if len(*p) > i+3 && i >= 0 {
		c := color.RGBAModel.Convert(c)
		rgba := c.(color.RGBA)
		(*p)[i] = byte(rgba.B)
		(*p)[i+1] = byte(rgba.G)
		(*p)[i+2] = byte(rgba.R)
		(*p)[i+3] = byte(rgba.A)
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	win, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
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
			p.Set(x, y, color.Gray{0})
		}
	}

	p1 := paddle{entity: entity{pos{x: 50, y: 100}, color.White}, w: 20, h: 100}
	p2 := paddle{entity: entity{pos{x: winWidth - 50, y: 200}, color.White}, w: 20, h: 100}
	b := ball{entity{pos{300, 300, 1, 1}, color.White}, 20}
	keyState := sdl.GetKeyboardState()
	frameTicker := time.NewTicker(5 * time.Millisecond)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}

		}
		p.clear()
		b.update(&p1, &p2)
		p1.update(keyState)
		p2.aiUpdate(&b)
		p1.draw(p)
		p2.draw(p)
		b.draw(p)

		tex.Update(nil, p, int(winWidth*int32(4)))
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		<-frameTicker.C
	}
}
