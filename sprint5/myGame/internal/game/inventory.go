package game

import "fmt"

type Inventory struct {
	Items map[string]int
}

// Take добавляет предмет в инвентарь
func (inv Inventory) Take(item string) string {
	if _, ok := inv.Items[item]; !ok {
		inv.Items[item]++
		return fmt.Sprintf("Вы положили %s в инвентарь", item)
	}
	inv.Items[item]++
	return fmt.Sprintf("Вы положили %s в инвентарь. Количество: %d", item, inv.Items[item])
}

// Drop удаляет предмет из инвентаря
func (inv Inventory) Drop(item string) string {
	if _, ok := inv.Items[item]; !ok {
		return fmt.Sprintf("У вас нет предмета %s", item)
	}
	if inv.Items[item] == 1 {
		delete(inv.Items, item)
		return fmt.Sprintf("Вы выбросили %s", item)
	}
	inv.Items[item]--
	return fmt.Sprintf("Вы выбросили %s\nОсталось %d", item, inv.Items[item])
}
