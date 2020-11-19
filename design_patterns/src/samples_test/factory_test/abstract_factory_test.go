package factory

import "testing"

import "../../samples/factory/abstract"

func TestAbstractFactory(t *testing.T) {
	factory := &abstract.SpringSkinFactory{}
	button := factory.CreateButton()
	button.Show()
	box := factory.CreateComboBox()
	box.Show()
}
