package main

import (
	"fmt"
)

func do(in, out chan int) {
	// добавьте ваш код здесь
	defer close(out)
	for v := range in {
		out <- 2 * v
	}
}

func main() {
	// создайте каналы
	chIn := make(chan int)
	chOut := make(chan int)

	// запустите горутину, которая преобразует числа
	go do(chIn, chOut)

	// запустите горутину, которая отправляет числа на обработку
	go func() {
		defer close(chIn)
		for i := 0; i <= 50; i++ {
			chIn <- i
		}
	}()

	// в цикле читайте числа из результирующего канала. Используйте for range
	for v := range chOut {
		fmt.Print(v, " ")
	}
}
