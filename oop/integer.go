package main

type Integer int

func (i Integer) Add(b Integer) {
	i += b
}

func main() {
	var a Integer = 1
	a.Add(2)
	print("a = ", a)
}
