package main

import "fmt"

type Generator struct {
	seed       int // seed нашего генератора
	prevNumber int // храним предыдущее число, выданное генератором
}

func (g *Generator) Int() int {
	// вычисляем и сохраняем новое псевдослучайное число
	g.prevNumber = (3*g.prevNumber + g.seed) % 11
	return g.prevNumber
}

func main() {
	// создаем гегератор с seed и выводим первые 3 число последовательности
	gen := Generator{seed: 12341}
	fmt.Println(gen.Int())
	fmt.Println(gen.Int())
	fmt.Println(gen.Int())

	// создаём генератор с seed=12 и выводим первые 3 числа последовательности
	gen2 := Generator{seed: 12}
	fmt.Println(gen2.Int())
	fmt.Println(gen2.Int())
	fmt.Println(gen2.Int())
}
