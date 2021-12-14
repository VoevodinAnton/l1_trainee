package tasks

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	v := make([]byte, size)
	fmt.Println("cap(v) = ", cap(v))

	for i := 0; i < size; i++ {
		v[i] = byte('a')
	}
	return string(v)
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println("\tstring", v)
	justString = v[:100]
	fmt.Println("\tjustString", justString)
}

func Run15() {
	someFunc()
}
