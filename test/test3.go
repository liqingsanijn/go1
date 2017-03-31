package main

import (
	"fmt"
)

func main() {
	const (
		a= iota
		b
		c= 0
		d= iota
	)
	fmt.Println("test3")
	fmt.Println(a, b, c, d)
	test3()
	e, f := test4()
	fmt.Println(e, f)
}

func test3()  {
	fmt.Println("函数test3")
}

func test4() (int, int) {
	fmt.Println("函数test4")
	return 1, 2
}
