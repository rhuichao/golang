package main

import (
	"fmt"
)

type Base struct {
	Name string
}

func (b *Base) Foo() {
	fmt.Println("Base Foo")
}

func (b *Base) Bar() {
	fmt.Println("Base Bar")
}

type Foo struct {
	Base
}

func (f *Foo) Foo() {
	f.Base.Foo()
	fmt.Println("Foo Foo")
}

func main() {
	foo := new(Foo)
	foo.Foo()
	foo.Bar()
}
