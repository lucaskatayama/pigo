package fonts

import (
	"embed"
	"fmt"
	"io"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed fonts/**
	fs embed.FS
)

var Fonts = map[string]font.Face{}

func init() {
	// https://www.pixelconverter.com/pixels-to-dpi-converter/ 128x64 0.96
	// LoadFont("DankMono", `/home/lucaskatayama/.fonts/Dank Mono Regular Nerd Font Complete.ttf`)
	LoadFont("pixelmix", `fonts/pixelmix.ttf`, 6, 96)
	// LoadFont("Grand9k", `/home/lucaskatayama/.fonts/Grand9K Pixel.ttf`, 6, 96)
	// LoadFont("pixelmix", `/home/lucaskatayama/.fonts/5x5_pixel.ttf`, 5, 149.07 )

}

func LoadFont(name string, path string, size float64, dpi float64) font.Face {
	//
	fmt.Println(name, path)
	if f, ok := Fonts[name]; ok {
		log.Println("font already loaded")
		return f
	}

	fi, err := fs.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	b, err := io.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}
	f, err := opentype.Parse(b)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}

	Fonts[name] = face

	return face
}
