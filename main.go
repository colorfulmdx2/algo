package main

import (
	"fmt"
)

// Что выведет следующая программа и сколько она будет выполняться по времени??

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	var foo []int
	var bar []int

	foo = append(foo, 1)
	foo = append(foo, 2)
	foo = append(foo, 3)
	bar = append(foo, 4)
	foo = append(foo, 5)

	fmt.Println(foo, bar)
}
