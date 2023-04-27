package menu

import (
	"sync"

	"github.com/lucaskatayama/pigo/internal/core/internet"
	"github.com/lucaskatayama/pigo/internal/core/speedtest"
	"github.com/lucaskatayama/pigo/internal/infra/icons"
)

var (
	_draw Feature = DisplayMenu
	mu    sync.RWMutex
)

func Draw() Feature {
	mu.RLock()
	r := _draw
	mu.RUnlock()
	return r
}

func SetDraw(f Feature) {
	mu.Lock()
	_draw = f
	mu.Unlock()
}

var DisplayMenu = &Menu{Items: []Item{}, Sel: 0}

func init() {
	a := []Item{
		{
			Icon:  icons.IconUpload,
			Label: "Upload",
			Sub:   &FeatureMock{},
		},
		{
			Icon:  icons.IconSpeed,
			Label: "SpeedTest",
			Sub: speedtest.New(func() {
				SetDraw(DisplayMenu)
			}),
		},
		{
			Icon:  icons.IconWifi,
			Label: "Internet",
			Sub: internet.New(func() {
				SetDraw(DisplayMenu)
			}),
		},
		{
			Icon:  icons.IconSD,
			Label: "Raspberry Pi",
			Sub:   &FeatureMock{},
		},
		{
			Icon:  icons.IconTelegram,
			Label: "Telegram",
			Sub:   &FeatureMock{},
		},
	}
	DisplayMenu.Items = append(DisplayMenu.Items, a...)
	DisplayMenu._len = len(a)
}
