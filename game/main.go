package main

import "fmt"

type Hero struct {
	Name   string
	Health int
	Damage int
	Def    int
}

// метод атаки
func (h Hero) Attack() {
	fmt.Println(h.Name, "нанес урон", h.Damage)
}

// блок
func (h Hero) Defense() {
	fmt.Println(h.Name, "заблокировал", h.Def, "единиц урона")
}

// ульта (в нашем случае лечит)
func (h Hero) Special(healthPoints int) {
	fmt.Println("Количество здоровья было", h.Health)
	h.Health += healthPoints
	fmt.Println("Количество здоровья стало", h.Health)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100, Damage: 30, Def: 20}
	myHero.Attack()
	myHero.Defense()
}
