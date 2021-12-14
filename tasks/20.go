package tasks

import (
	"fmt"
	"strings"
)

var sentence = "snow dog sun"

func reverseSentence(sentence string) (reverseSentence string) {
	words := strings.Split(sentence, " ")
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

func Run20() {
	fmt.Println("source sentence =", sentence)

	reverseSentence := reverseSentence(sentence)

	fmt.Println("reverse sentence =", reverseSentence)
}
