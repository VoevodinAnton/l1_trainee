package tasks

import (
	"fmt"
)

var (
	number   int64 = 2
	bitPos   int   = 1
	bitValue int   = 0
)

func SetBit(number int64, bitPos int, bitValue int) int64 {
	if bitPos > 64 {
		fmt.Println("The number of bits cannot be more than 64")
		return 0
	}

	if int((number&(1<<bitPos))>>bitPos) != bitValue {
		if bitValue == 1 {
			number |= 1 << bitPos
		} else {
			number &^= 1 << bitPos
		}
	}
	return number
}

func Run8() {
	number = SetBit(number, bitPos, bitValue)
	fmt.Printf("bin(%v) = %b", number, number)
}
