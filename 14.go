package main

import (
	"fmt"
)

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

func printInventory(inventory []InventoryItem) {
	for i, item := range inventory {
		questMark := ""
		if item.IsQuestItem {
			questMark = " [Квестовый]"
		}
		fmt.Printf("%d. %s - %.1f кг%s\n", i+1, item.Name, item.Weight, questMark)
	}
}

func calculateTotalWeight(inventory []InventoryItem) float64 {
	total := 0.0
	for _, item := range inventory {
		total += item.Weight
	}
	return total
}

func main() {
	inventory := []InventoryItem{}
	
	inventory = append(inventory, InventoryItem{"Меч", 5.2, false})
	inventory = append(inventory, InventoryItem{"Щит", 8.1, false})
	inventory = append(inventory, InventoryItem{"Зелье здоровья", 0.5, false})
	inventory = append(inventory, InventoryItem{"Древний артефакт", 3.7, true})
	inventory = append(inventory, InventoryItem{"Золотые монеты", 2.3, false})
	
	fmt.Println("Инвентарь персонажа:")
	printInventory(inventory)
	
	totalWeight := calculateTotalWeight(inventory)
	fmt.Printf("\nОбщий вес инвентаря: %.1f кг\n", totalWeight)
	
	inventory = append(inventory, InventoryItem{"Лук", 3.2, false})
	inventory = append(inventory, InventoryItem{"Ключ от темницы", 0.2, true})
	
	fmt.Println("Инвентарь с новыми вещами:")
	printInventory(inventory)
	
	totalWeight = calculateTotalWeight(inventory)
	fmt.Printf("Вес после добавления новых предметов: %.1f кг\n", totalWeight)

	fmt.Println("Квестовые предметы:")
	for _, item := range inventory {
		if item.IsQuestItem {
			fmt.Printf("- %s (%.1f кг)\n", item.Name, item.Weight)
		}
	}
}