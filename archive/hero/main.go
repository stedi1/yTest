package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name       string
	BaseDamage int
	BaseDef    int
	MinDamage  int
	MaxDamage  int
	MinDef     int
	MaxDef     int
}

// возвращает строку с нанесенным уроном
func (h Hero) Attack() string {
	return fmt.Sprint(h.Name, " нанес урон, равный ", h.BaseDamage+randInt(h.MinDamage, h.MaxDamage))
}

// возвращает строку с заблокированным уроном
func (h Hero) Defense() string {
	return fmt.Sprint(h.Name, " заблокировал урон, равный ", h.BaseDef+randInt(h.MinDef, h.MaxDef))
}

func main() {
	// создаем структуру (Героя Артура)
	arthur := Hero{Name: "Артур", BaseDamage: 10, BaseDef: 7,
		MinDamage: 5, MaxDamage: 20, MinDef: 3, MaxDef: 15}
	// юзаем методы Артура
	attackArthur := arthur.Attack()
	fmt.Println(attackArthur)
	defenceArthur := arthur.Defense()
	fmt.Println(defenceArthur)

	// еще герой Бернал
	bernal := Hero{Name: "Бернал", BaseDamage: 12, BaseDef: 10,
		MinDamage: 7, MaxDamage: 25, MinDef: 5, MaxDef: 20}
	// юзаем его методы
	attackBernal := bernal.Attack()
	fmt.Println(attackBernal)
	defenseBernal := bernal.Defense()
	fmt.Println(defenseBernal)
}

// randInt возвращает случайное значение из диапазона [min, max]
func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
