package tasks

import "fmt"

type Human struct {
	FName string
	LName string
	Age   int
}

type Action struct {
	Human Human
}

func (a *Action) Talk() {
	fmt.Print("Hi, my name is ", a.Human.FName, ". I'm 21 years old.")
}

func Run1() {
	a := Action{}
	a.Human = Human{FName: "Anton", LName: "Voevodin", Age: 21}

	a.Talk()
}
