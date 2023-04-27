package control

type Control interface {
	OnEnter()
	OnClick()
	OnDoubleClick()
	OnPress()
}
