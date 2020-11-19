package abstract

type SkinFactory interface {
	CreateButton() Button
	CreateComboBox() ComboBox
}

type SummerSkinFactory struct {
}

type SpringSkinFactory struct {
}

func (*SpringSkinFactory) CreateButton() Button {
	return &SpringButton{}
}

func (*SpringSkinFactory) CreateComboBox() ComboBox {
	return &SpringComboBox{}
}

func (*SummerSkinFactory) CreateButton() Button {
	return &SummerButton{}
}

func (*SummerSkinFactory) CreateComboBox() ComboBox {
	return &SummerComboBox{}
}
