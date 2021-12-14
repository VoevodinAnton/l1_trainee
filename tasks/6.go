package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker6(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(10)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func Run6() {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker6(ctx, i, result)
	}

	foundBy := <-result
	fmt.Println("result found by", foundBy)

	finish()

	time.Sleep(time.Second)
}
