package tasks

import "fmt"

func fillIndicatorMap(indicatorMap *map[int]bool, slice []int) {
	for _, v := range slice {
		if _, ok := (*indicatorMap)[v]; ok {
			(*indicatorMap)[v] = true
		} else {
			(*indicatorMap)[v] = false
		}
	}
}

func intersect(slice1 []int, slice2 []int) []int {
	indicatorMap := make(map[int]bool)

	fillIndicatorMap(&indicatorMap, slice1)
	fillIndicatorMap(&indicatorMap, slice2)

	var intersection []int
	for k, v := range indicatorMap {
		if v {
			intersection = append(intersection, k)
		}
	}
	return intersection
}

func Run11() {
	a := []int{2, 4, 6, 8, 10}
	b := []int{1, 3, 4, 6, 5}

	c := intersect(a, b)

	fmt.Println("sliceA", a)
	fmt.Println("sliceB", b)
	fmt.Println("intersection of the sets", c)
}
