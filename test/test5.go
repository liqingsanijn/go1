package main

import (
	"fmt"
)

func testCallback(x int, callback func(param int) int) int {
	return callback(x)
}
func main()  {
	fmt.Println(testCallback(1, addOne))
	fmt.Println(testCallback(2, func(param int) int {
		param +=2
		return param
	}))
}

func addOne(x int) int {
	x++
	return x
}