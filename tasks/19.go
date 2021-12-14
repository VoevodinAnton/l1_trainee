package tasks

var word = "abcde"

func reverseWord(str string) (s string) {
	word := []rune(str)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}

func Run19() {
	println("source string =", word)

	reverseString := reverseWord(word)

	println("reversed string =", reverseString)
}
