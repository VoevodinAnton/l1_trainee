package tasks

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	Number int32
}

func (counter *Counter) add() {
	atomic.AddInt32(&counter.Number, 1)
}

func Run18() {
	var counter Counter

	for i := 0; i < 100; i++ {
		go counter.add()
	}

	time.Sleep(time.Second)
	fmt.Println("\tCounter", counter)
}
