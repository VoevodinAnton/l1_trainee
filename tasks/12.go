package tasks

import "fmt"

func doSubset(data []string) []string {
	indicatorMap := make(map[string]bool)

	for _, v := range data {
		indicatorMap[v] = true
	}

	var subset []string
	for k, v := range indicatorMap {
		if v {
			subset = append(subset, k)
		}
	}
	return subset
}

func Run12() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	subset := doSubset(set)

	fmt.Println("set", set)
	fmt.Println("subset", subset)
}
