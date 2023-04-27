package menu

import (
	"fmt"

	"github.com/lucaskatayama/pigo/internal/infra/icons"
	"github.com/lucaskatayama/pigo/internal/ui/control"
	"github.com/lucaskatayama/pigo/internal/ui/display"
)

type Feature interface {
	control.Control
	Draw(s display.Display)
}

type Item struct {
	Icon  icons.ImageBytes
	Label string
	Sub   Feature
}

type FeatureMock struct{}

func (f *FeatureMock) Draw(s display.Display) {
	s.Begin()
	s.Text("alskdnalsd", 0, 0)
	s.Flush()
}

func (f *FeatureMock) OnClick() {
	fmt.Println("ajsdja")
}

func (f *FeatureMock) OnDoubleClick() {
	SetDraw(DisplayMenu)
}

func (f *FeatureMock) OnPress() {

}

func (f *FeatureMock) OnEnter() {}
