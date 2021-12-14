package tasks

import "fmt"

var words = []string{
	"asd",
	"фыв",
	"aasd",
	"фывв",
	"世界",
	"世界界",
	" ",
	"  ",
}

func isUnique(s string) bool {
	runes := []rune(s)
	m := make(map[rune]bool)
	for _, rune := range runes {
		if _, ok := m[rune]; ok {
			return false
		} else {
			m[rune] = true
		}
	}
	return true
}

func Run26() {
	for _, s := range words {
		fmt.Printf("All symbols of (%v) are uniq: %v\n", s, isUnique(s))
	}
}
