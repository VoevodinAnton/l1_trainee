package tasks

import "fmt"

func binarySearch(sortedList []int, lookingFor int) int {
	var lo int = 0
	var hi int = len(sortedList) - 1

	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue int = sortedList[mid]

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			// We want to use the left half of our list
			hi = mid - 1
		} else {
			// We want to use the right half of our list
			lo = mid + 1
		}
	}
	return -1
}

func Run17() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 2))
}
