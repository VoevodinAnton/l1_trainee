package tasks

import "fmt"

func Run13() {
	a := "Лупа"
	b := "Пупа"

	fmt.Print("a = ", a, ", b = ", b, "\n")

	a, b = b, a

	fmt.Print("a = ", a, ", b = ", b)
}
