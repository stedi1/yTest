package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// в качестве seed передаем методу текущее время в unix формате (в виде числа)
	src := rand.NewSource(time.Now().Unix())
	for i := 0; i < 10; i++ {
		randomNumber := src.Int63()
		fmt.Println(randomNumber%6 + 1)
	}
}
