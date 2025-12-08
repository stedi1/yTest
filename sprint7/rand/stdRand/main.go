package main

import (
	"fmt"
	"math/rand"
)

func main() {
	src := rand.NewSource(123)
	for i := 0; i < 10; i++ {
		// получаем случайное число
		randomNumber := src.Int63()
		// так как число варьируется от 0 до 2^63, то мы берём остаток
		// от деления randomNumber на 6, чтобы получить число от 0 до 5.
		// Дальше прибавляем 1 и получаем число от 1 до 6
		fmt.Println(randomNumber%6 + 1)
	}
}
