package abstract

import "fmt"

type Button interface {
	Show()
}

type ComboBox interface {
	Show()
}

type SpringButton struct {
}

type SpringComboBox struct {
}

type SummerButton struct {
}

type SummerComboBox struct {
}

func (*SpringButton) Show() {
	fmt.Println("This is a spring button!")
}

func (*SpringComboBox) Show() {
	fmt.Println("This is a spring combo box!")
}

func (*SummerButton) Show() {
	fmt.Println("This is a summer button")
}

func (*SummerComboBox) Show() {
	fmt.Println("This is a summer combo box")
}
