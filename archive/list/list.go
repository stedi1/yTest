// list.go
package list

import (
	"fmt"
)

func Min(list []int) int {
	var min int
	if len(list) < 1 {
		return 0
	}
	min = list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}

func Mul(a, b []int) ([]int, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("different length of slices")
	}
	result := make([]int, len(a))
	for i, v := range a {
		result[i] = v * b[i]
	}
	return result, nil
}

// напишите рекурсивную функцию GCD, которая вычисляет НОД
// двух чисел по алгоритму Евклида
func GCD(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}
	r := a % b
	if r == 0 {
		return b
	}
	return GCD(b, r)
}
