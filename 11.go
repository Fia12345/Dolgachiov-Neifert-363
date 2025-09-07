package main
import "fmt"

type Product struct {
	Name string
	Category string
	Price float64
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
	var filtered []Product
	
	for _, product := range products {
		if product.Price <= maxPrice && product.Category == category {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

func main() {
	products := []Product{
		{"МакБук", "Электроника", 80000},
		{"Эйфон", "Электроника", 50000},
		{"Книга 1984 Джорджа Оруэлла", "Разное", 500},
		{"ЭйПэд", "Электроника", 45000},
		{"ЭйРучка", "Канцелярия", 50},
		{"Монитор Эппле Студио Дисплей Серебристый", "Электроника", 150000},
		{"Блокнот", "Канцелярия", 200},
		{"Носинг Ийр", "Электроника", 5000},
	}
	
	filtered := filterProducts(products, 30000, "Электроника")
	
	fmt.Println("Все товары:")
	for i, product := range products {
		fmt.Printf("%d. %s (%s) - %.2f руб.\n", i+1, product.Name, product.Category, product.Price)
	}
	
	fmt.Println("\nОтфильтрованные товары (Электроника до 30000 руб.):")
	for i, product := range filtered {
		fmt.Printf("%d. %s - %.2f руб.\n", i+1, product.Name, product.Price)
	}
	
	fmt.Println("\nКанцелярия до 100 руб.:")
	officeFiltered := filterProducts(products, 100, "Канцелярия")
	for i, product := range officeFiltered {
		fmt.Printf("%d. %s - %.2f руб.\n", i+1, product.Name, product.Price)
	}
}