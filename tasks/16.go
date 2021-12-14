package tasks

import "fmt"

var (
	lst = []int{5, 4, 1, 2, 0, -1, 1234, 32, 12, 2345, 99, -8, 45}
)

func qsort(lst []int, left int, right int) []int {

	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := lst[(left+right)/2]

	//Цикл, начинающий саму сортировку
	for l <= r {

		//Ищем значения больше 'центра'
		for lst[r] > center {
			r--
		}

		//Ищем значения меньше 'центра'
		for lst[l] < center {
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			lst[r], lst[l] = lst[l], lst[r]
			l++
			r--
		}
	}

	if r > left {
		qsort(lst, left, r)
	}

	if l < right {
		qsort(lst, l, right)
	}

	return lst
}

func Run16() {
	fmt.Println("source slice", lst)
	lst = qsort(lst, 0, len(lst)-1)
	fmt.Println("sorted slice", lst)
}
