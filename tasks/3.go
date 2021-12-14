package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run3() {
	numbers := []int{2, 4, 6, 8, 9, 10}
	var result float64

	start := time.Now()
	for i := 0; i < len(numbers); i++ {
		go sumOfSquares(numbers, i, &result)
	}
	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
	fmt.Println("Result is ", result)
}

func sumOfSquares(numbers []int, id int, result *float64) {
	*result = *result + math.Pow(float64(numbers[id]), 2)
}
