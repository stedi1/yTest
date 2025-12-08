package main

import (
	"fmt"
	"math/rand"
)

func Winner(rnd *rand.Rand, bets map[int]string) string {
	var minDif int64 = 100
	winner := ""
	number := rnd.Int63() % 100

	for num, name := range bets {
		tmpDif := number - int64(num)
		if tmpDif < 0 {
			tmpDif = -tmpDif
		}
		if tmpDif < minDif {
			minDif = tmpDif
			winner = name
		}
	}
	return winner
}

func main() {
	rnd := rand.New(rand.NewSource(1001))
	// кто какое число загадал
	bets := map[int]string{
		20: "Маша",
		34: "Игорь",
		77: "Олег",
		51: "Света",
		2:  "Саша",
	}

	fmt.Println("Победитель:", Winner(rnd, bets))
}
