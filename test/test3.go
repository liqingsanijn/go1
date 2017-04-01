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
	x,y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
	var t Test3
	t.x = 5
	t.y = 10
	fmt.Println(t.multi())
}

func test3()  {
	fmt.Println("函数test3")
}

func test4() (int, int) {
	fmt.Println("函数test4")
	return 1, 2
}

func swap(x *int, y *int)  {
	temp := *x
	*x = *y
	*y = temp
}

type Test3 struct {
	x int
	y int
}

func (t Test3) multi(t1 Test3) int {
	return t.x*t.y
}
