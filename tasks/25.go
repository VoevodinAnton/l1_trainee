package tasks

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
}

func Run25() {
	fmt.Println("Привет")
	Sleep(time.Second * 3)
	fmt.Println("Пока")
}
