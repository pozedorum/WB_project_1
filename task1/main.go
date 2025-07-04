package main

import "fmt"

type Human struct {
	age   int
	name  string
	alive bool
}

func (h *Human) GetAge() int {
	return h.age
}

func (h *Human) IsItMyName(name2 string) bool {
	return h.name == name2
}

type Action struct {
	Human
	NewProperty string
}

func MakeAction(age int, name, NewProperty string, alive bool) *Action {
	return &Action{
		Human: Human{
			age:   age,
			name:  name,
			alive: alive,
		},
		NewProperty: NewProperty,
	}
}

func main() {
	newAction := MakeAction(12, "aaa", "bbb", false)
	if newAction.IsItMyName("aa") {
		fmt.Println("YES, it's my name")
	} else {
		fmt.Println("no no no mister fish")
	}
}
