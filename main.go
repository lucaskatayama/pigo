package main

import (
	"log"
	"time"

	"github.com/lucaskatayama/pigo/internal/infra/onebutton"
	"github.com/lucaskatayama/pigo/internal/infra/ssd1306"
	"github.com/lucaskatayama/pigo/internal/ui/menu"
	"golang.org/x/image/font/basicfont"

	"periph.io/x/host/v3"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	s, err := ssd1306.New()
	if err != nil {
		log.Fatal(err)
	}

	btn, err := onebutton.New("26")
	if err != nil {
		log.Fatal(err)
	}

	s.SetDefaultFont(basicfont.Face7x13)

	btn.SetClick(func() {
		menu.Draw().OnClick()
	})

	btn.SetDoubleClick(func() {
		menu.Draw().OnDoubleClick()
	})

	btn.SetLongPress(func() {
		menu.Draw().OnPress()
	})

	for {
		menu.Draw().Draw(s)
		<-time.After(50 * time.Millisecond)
	}

}
