package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run2() {
	numbers := []int{2, 4, 6, 8, 9, 10}

	start := time.Now()
	for i := 0; i < len(numbers); i++ {
		go doOperation(numbers, i)
	}
	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}

func doOperation(numbers []int, id int) {
	fmt.Println(math.Pow(float64(numbers[id]), 2))
}
