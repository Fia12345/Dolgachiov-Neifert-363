package main

import (
	"fmt"
)

func printExpenses(expenses map[string]float64) {
	for category, amount := range expenses {
		fmt.Printf("- %s: %.2f руб.\n", category, amount)
	}
}

func calculateTotal(expenses map[string]float64) float64 {
	total := 0.0
	for _, amount := range expenses {
		total += amount
	}
	return total
}

func main() {
	expenses := make(map[string]float64)
	
	expenses["Еда"] = 15000
	expenses["Транспорт"] = 5000
	expenses["Развлечения"] = 3000
	
	fmt.Println("Начальные траты:")
	printExpenses(expenses)
	
	expenses["Еда"] += 2000
	fmt.Println("\nДобавление:")
	printExpenses(expenses)
	
	total := calculateTotal(expenses)
	fmt.Printf("\nИтоговая сумма: %.2f руб.\n", total)
	
	total = calculateTotal(expenses)
	fmt.Printf("Общая сумма после новых категорий: %.2f руб.\n", total)
}