package main

import (
	"fmt"
	"mygame/internal/game"
)

func main() {
	// Общая структура для всех классов с полями «Здоровье», «Урон» и «Защита»
	defaultHero := game.Hero{Health: 100, Damage: 30, Def: 20}

	// Маг
	// Инвентарь для мага
	mageInventory := game.Inventory{Items: make(map[string]int)}
	// Структура для мага
	mage := game.Mage{Hero: defaultHero, Inventory: mageInventory, Name: "Мерлин", ClassName: "Маг"}
	// Представимся
	fmt.Println("Я", mage.Name, "класса", mage.ClassName)
	fmt.Println(mage.Attack())
	fmt.Println(mage.Defense())
	fmt.Println(mage.Buff())
	fmt.Println(mage.Attack())
	// Маг кладёт посох в инвентарь
	fmt.Println(mage.Take("Посох"))
	// Посох не понравился, выбрасывает посох
	fmt.Println(mage.Inventory.Drop("Посох"))
	fmt.Println()

	// Воин
	warriorInventory := game.Inventory{Items: make(map[string]int)}
	warrior := game.Warrior{Hero: defaultHero, Inventory: warriorInventory, Name: "Арагорн", ClassName: "Воин"}
	fmt.Println("Я", warrior.Name, "класса", warrior.ClassName)
	fmt.Println(warrior.Attack())
	fmt.Println(warrior.Defense())
	fmt.Println(warrior.Buff())
	// Воин кладёт в инвентарь меч
	fmt.Println(warrior.Take("Меч"))
	// Воин кладёт в инвентарь шлем
	fmt.Println(warrior.Take("Шлем"))
	// Воин кладёт в инвентарь наручи
	fmt.Println(warrior.Take("Наручи"))
	// Ещё один шлем в инвентарь
	fmt.Println(warrior.Take("Шлем"))
	// Много шлемов, один выкинул
	fmt.Println(warrior.Drop("Шлем"))
	// Попробовал выкинуть сапоги, но их же и так нет
	fmt.Println(warrior.Drop("Сапоги"))
	fmt.Println()

	// Лекарь
	healerInventory := game.Inventory{Items: make(map[string]int)}
	healer := game.Healer{Hero: defaultHero, Inventory: healerInventory, Name: "Елена Малышева", ClassName: "Лекарь"}
	fmt.Println("Я", healer.Name, "класса", healer.ClassName)
	fmt.Println(healer.Attack())
	fmt.Println(healer.Defense())
	fmt.Println(healer.Buff())
	// Лекарь кладёт в инвентарь амулет
	fmt.Println(healer.Take("Амулет"))
	// Лекарь кладёт в инвентарь плащ
	fmt.Println(healer.Take("Плащ"))
}
