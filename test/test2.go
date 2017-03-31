package main
import (
	"fmt"
)

func main() {
	var a bool
	fmt.Println(a)
	testMap := make(map[string]string)
	testMap["key1"] = "value1"
	testMap["key2"] = "value2"
	for key,value := range testMap {
		fmt.Println("key:" + key + ",value:" + value)
	}
	arr := [...]int{1, 2, 3}
	for i,v := range arr {
		fmt.Printf("i:%d,v:%d\n", i, v)
	}
	fmt.Println("a", "b", "c")
}
