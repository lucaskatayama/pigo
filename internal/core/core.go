package core

type FeatureI struct{}

func (f *FeatureI) OnEnter()       {}
func (f *FeatureI) OnClick()       {}
func (f *FeatureI) OnDoubleClick() {}
func (f *FeatureI) OnPress()       {}
