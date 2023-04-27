package upload

import (
	"github.com/lucaskatayama/pigo/internal/core"
	"github.com/lucaskatayama/pigo/internal/ui/display"
)

type Upload struct {
	core.FeatureI
}

func New() *Upload {
	return &Upload{}
}

func (u *Upload) Draw(d display.Display) {
	d.Begin()
	d.Text("Auto Upload", 2, 2)

	d.Flush()
}

func (u *Upload) OnEnter() {

}

func (u *Upload) OnClick() {

}
