package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))

	human := Human{"Sam", 22, "135464e453xxx"}
	student := Student{Human{"Mike", 25, "1585810xxxx"}, "csu", 100}
	x := make([]Men, 2)
	x[0], x[1] = human, student

	for _, value := range x {
		value.SayHi()
	}
}


type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

func (h Human) SayHi()  {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) String() string {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
	return h.name
}

func (s Student) SayHi()  {
	fmt.Println("Hi, I am %s you can call me on %s\n", s.name, s.phone)
}

type Men interface {
	SayHi()
}