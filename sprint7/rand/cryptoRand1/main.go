package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 10; i++ {
		// получаем случайное число. Укажем максимальное значение 6, чтобы получать числа от 0 до 5.
		// rand.Reader - криптографически стойкий генератор, который дает нам Go
		randNumber, err := rand.Int(rand.Reader, big.NewInt(6))
		// не забываем проверить ошибку на nil
		if err != nil {
			panic(err)
		}
		// прибавляем 1, чтобы получить числа от 1 до 6.
		fmt.Println(randNumber.Int64() + 1)
	}
}
