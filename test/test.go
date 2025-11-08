package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}
	sum := 0
	for i, v := range nums {
		if i%2 == 0 {
			sum += v
		}
	}
	fmt.Println(sum)

	// указатели
	fmt.Println("\nУказатели")
	var i, j int
	pi := &i
	pj := &j
	*pi = 7
	*pj = *pi
	*pj += 3
	px := pi
	*px = 5
	fmt.Println(i, j, px, "\n")

	// работа с указателями (заполнить массив указателями, изменить значения используя указатели)
	// исходный массив
	numbers := [...]int{1, 3, 8, 19, 42}

	// 1. Создайте и заполните слайс указателей на элементы массива
	// 2. Пройдитесь по слайсу и умножьте на три все значения, на которые
	//    ссылаются указатели
	ptrs := make([]*int, len(numbers))
	for i := range numbers {
		ptrs[i] = &numbers[i]
	}
	for i := range ptrs {
		*ptrs[i] *= 3
	}

	fmt.Println(numbers)
}
