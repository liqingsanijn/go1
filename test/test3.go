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
}
