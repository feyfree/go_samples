package snippets

// 结构体嵌套接口 接口没有变量名是可以通过编译的
// 但是在初始化的时候， 如果需要复制的话,
// 需要增加一个该接口实现的一个实例作为构造参数

import (
	"fmt"
	"testing"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

type student struct {
	name string
}

func (s student) work() {
	fmt.Println("My name is " + s.name)
}
func TestInterfaceWithStruct(t *testing.T) {
	w := person{}
	fmt.Println(w)
	m := person{worker: student{"Lei"}}
	fmt.Println(m)
}
