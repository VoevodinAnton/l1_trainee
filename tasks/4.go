package tasks

import (
	"fmt"
)

const workerCount = 3

func worker4(workerNum int, in <-chan interface{}) {
	for i := range in {
		fmt.Println("\tThe worker", workerNum, "got the value", i)
	}
}

func Run4() {
	in := make(chan interface{})

	numbers := []interface{}{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, "dog", "cat",
		struct {
			FName string
			Lname string
			Age   int
		}{"Anton", "Voevodin", 21}}

	for i := 0; i < workerCount; i++ {
		go worker4(i, in)
	}

	for _, value := range numbers {
		in <- value
	}
	close(in)

	fmt.Scanln()
}
