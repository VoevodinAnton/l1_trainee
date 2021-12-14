package tasks

import "fmt"

func getDecade(num float32) int {
	if num > 0 {
		return int(num/10) * 10
	} else {
		return int(num/10-1) * 10
	}
}

func Run10() {
	input := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 123.3, 3, -3}

	decadeGroups := make(map[int][]float32)

	for _, v := range input {
		decade := getDecade(v)
		decadeGroups[decade] = append(decadeGroups[decade], v)
	}
	fmt.Println(decadeGroups)
}
