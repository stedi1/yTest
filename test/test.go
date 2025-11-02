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
}
