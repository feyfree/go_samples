package prototype

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Shape interface {
	Show()
}

type Circle struct {
	name string
}

func (circle *Circle) Show() {
	fmt.Println(circle.name)
}
