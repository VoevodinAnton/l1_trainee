package tasks

import (
	"fmt"
	"time"
)

func Run5() {
	in := make(chan int)

	go func(out chan<- int) {
		var i int
		for {
			out <- i
			i++
		}
	}(in)

	go func(in <-chan int) {
		for v := range in {
			fmt.Println(v)
		}
	}(in)

	time.Sleep(time.Second)
	close(in)
}
