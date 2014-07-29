package main

import (
	"fmt"
)

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main() {
	// rect := new(Rect)
	rect := &Rect{0, 0, 100, 200}
	fmt.Println("rect1 area is ", rect.Area())
	fmt.Println("rect1 x is ", rect.x)
	fmt.Println("rect2 area is ", NewRect(0, 0, 10, 20).Area())
}
