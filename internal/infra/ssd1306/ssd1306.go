package ssd1306

import (
	"image"
	"image/color"

	"github.com/lucaskatayama/pigo/internal/infra/icons"
	"github.com/lucaskatayama/pigo/internal/ui/fonts"
	"golang.org/x/image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
)

type SSD1306Display struct {
	frame *image1bit.VerticalLSB
	font  *basicfont.Face

	Dev    *ssd1306.Dev
	Bounds image.Rectangle
}

func New() (*SSD1306Display, error) {
	bus, err := i2creg.Open("1")
	if err != nil {
		return nil, err
	}

	// Open a handle to a ssd1306 connected on the I²C bus:
	// dev, err := ssd1306.NewI2C(bus, 128, 64, false)
	dev, err := ssd1306.NewI2C(bus, &ssd1306.Opts{W: 128, H: 64, Rotated: false})
	if err != nil {
		return nil, err
	}

	return &SSD1306Display{Dev: dev, Bounds: dev.Bounds()}, nil
}

func (ssd *SSD1306Display) SetDefaultFont(f *basicfont.Face) {
	ssd.font = f
}

func (ssd *SSD1306Display) TextFont(s string, name string, x, y int) int {
	face := fonts.Fonts[name]

	drawer := font.Drawer{
		Dst:  ssd.frame,
		Src:  &image.Uniform{image1bit.On},
		Face: face,
		Dot:  fixed.P(x+ssd.font.Advance, y+ssd.font.Ascent),
	}

	drawer.DrawString(s)
	return face.Metrics().Ascent.Round()
}

func (ssd *SSD1306Display) Text(s string, x, y int) int {
	drawer := font.Drawer{
		Dst:  ssd.frame,
		Src:  &image.Uniform{image1bit.On},
		Face: ssd.font,
		Dot:  fixed.P(x+ssd.font.Advance, y+ssd.font.Ascent),
	}
	drawer.DrawString(s)
	return ssd.font.Ascent
}

func (ssd *SSD1306Display) Begin() {
	ssd.frame = image1bit.NewVerticalLSB(ssd.Dev.Bounds())
}

func (ssd *SSD1306Display) DrawImageBytes(i icons.ImageBytes, x, y int) {
	ssd.DrawBytes(i.B, x, y, i.Dx, i.Dy)
}

func (ssd *SSD1306Display) DrawRect(x, y, dx, dy int) {
	rect := image.Rect(x, y, x+dx, y+dy) //  geometry of 2nd rectangle which we draw atop above rectangle
	rectColor := color.RGBA{255, 255, 255, 255}

	// create a red rectangle atop the green surface
	// draw.Draw(ssd.frame, rect, &image.Uniform{rectColor}, image.ZP, draw.Src)
	draw.Src.Draw(ssd.frame, image.Rect(x, y, x+rect.Dx(), y+rect.Dy()), &image.Uniform{rectColor}, image.Point{})
}

func (ssd *SSD1306Display) DrawBytes(b []byte, x, y int, dx, dy int) {
	// src := image.NewAlpha(image.Rectangle{Max: image.Point{dx, dy}})
	src := image1bit.NewVerticalLSB(image.Rectangle{Max: image.Point{dx, dy}})

	copy(src.Pix, b)

	draw.Src.Draw(ssd.frame, image.Rect(x, y, x+src.Rect.Dx(), y+src.Rect.Dy()), src, image.Point{})
}

func (ssd *SSD1306Display) Flush() {
	ssd.Dev.Write(ssd.frame.Pix)
}
