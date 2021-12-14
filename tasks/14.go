package tasks

import (
	"fmt"
)

func Run14() {
	words := []interface{}{"cat", 12, false, 5, "table", make(chan string)}

	for _, v := range words {
		switch v.(type) {
		case int:
			fmt.Println("\t", v, "- тип int")
		case bool:
			fmt.Println("\t", v, "- тип bool")
		case string:
			fmt.Println("\t", v, "- тип string")
		case chan string:
			fmt.Println("\t", v, "- тип channel")
		}
	}
}
