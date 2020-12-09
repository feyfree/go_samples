package prototype

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Job interface {
	Work()
}

type Driver struct {
	Name string
}

func (driver *Driver) Work() {
	fmt.Println("My name is " + driver.Name + ". I am a driver")
}

type Teacher struct {
	Name string
}

func (driver *Driver) Clone() Cloneable {
	d := *driver
	return &d
}

func (teacher *Teacher) Clone() Cloneable {
	t := *teacher
	return &t
}

func (teacher *Teacher) Work() {
	fmt.Println("My name is " + teacher.Name + ". I am a teacher")
}
