package display

import (
	"github.com/lucaskatayama/pigo/internal/infra/icons"
	"golang.org/x/image/font/basicfont"
)

type Display interface {
	Begin()
	Flush()

	SetDefaultFont(f *basicfont.Face)
	TextFont(s string, name string, x, y int) int
	Text(s string, x, y int) int

	DrawImageBytes(i icons.ImageBytes, x, y int)
	DrawBytes(b []byte, x, y int, dx, dy int)

	DrawRect(x, y, dx, dy int)
}
