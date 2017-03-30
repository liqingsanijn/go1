package main
import "fmt"

func main() {
	fmt.Println("hello,world")
	a := 1
	fmt.Println(a)
	b := new(test)
 	b.test1 = 0
	fmt.Println(b.test1)	
}
type test struct {
	test1 int
	test2 int
}