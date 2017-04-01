package main

import (
	"fmt"
)

func main()  {
	var p Person = new(Programmer)
	p.setAge(20)
	p.setName("jx")
	fmt.Printf("name:%s,age:%d\n", p.getName(), p.getAge())
	slice := []int {0, 1, 2, 3 ,4}
	for i := 5; i < 10; i++ {
		slice = append(slice, i)
	}
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	newSlice[0] = 100
	fmt.Println(slice)
	fmt.Println(newSlice)
	testMap := make(map[string]string)
	testMap["key1"] = "value1"
	testMap["key2"] = "value2"
	fmt.Println(testMap);
	for key, value := range testMap {
		fmt.Printf("%s:%s\n", key, value)
	}
	delete(testMap, "key2")
	fmt.Println(testMap)

}

type Person interface {
	getAge() int
	getName() string
	setAge(age int)
	setName(name string)
}

type Programmer struct {
	age int
	name string
}

func (p Programmer) getAge() int {
	return p.age
}

func (p Programmer) getName() string {
	return p.name
}

func (p *Programmer) setAge(age int) {
	p.age = age
}

func (p *Programmer) setName(name string) {
	p.name = name
}
