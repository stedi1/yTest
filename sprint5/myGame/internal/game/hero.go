package game

import "fmt"

type Hero struct {
	Health int // здоровье
	Damage int // урон
	Def    int // защита
}

func (h Hero) Attack() string {
	return fmt.Sprint("Ваш персонаж нанёс урон, равный ", h.Damage)
}

func (h Hero) Defense() string {
	return fmt.Sprint("Ваш персонаж заблокировал ", h.Def, " урона")
}

type Warrior struct {
	Hero
	Inventory
	Name      string
	ClassName string
}

// Buff - спец. умение для Warrior
func (w *Warrior) Buff() string {
	w.Def += 20
	return fmt.Sprintf("%s класса %s увеличил свою защиту.\nЗащита %s теперь %d.",
		w.Name,
		w.ClassName,
		w.Name,
		w.Def,
	)
}

type Mage struct {
	Hero
	Inventory
	Name      string
	ClassName string
}

// Buff - спец. умение для Mage
func (m *Mage) Buff() string {
	m.Damage += 30
	return fmt.Sprintf("%s класса %s усилил свою атаку.\nАтака %s теперь %d.",
		m.Name,
		m.ClassName,
		m.Name,
		m.Damage,
	)
}

type Healer struct {
	Hero
	Inventory
	Name      string
	ClassName string
}

// Buff - спец. умение для Healer
func (h *Healer) Buff() string {
	h.Health += 50
	return fmt.Sprintf("%s класса %s увеличил своё здоровье.\nЗдоровье %s теперь %d.",
		h.Name,
		h.ClassName,
		h.Name,
		h.Health,
	)
}
