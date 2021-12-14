package tasks

import (
	"fmt"
	"math/big"
)

func Run22() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("12412341023859823457923523", 10)
	b.SetString("12412341023859524352432345", 10)

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Printf("Частное: %v\n", big.NewInt(0).Div(a, b))
	fmt.Printf("Произведение: %v\n", big.NewInt(0).Mul(a, b))
	fmt.Printf("Сумма: %v\n", big.NewInt(0).Add(a, b))
	fmt.Printf("Разность: %v\n", big.NewInt(0).Sub(a, b))
}
