package main
import "fmt"

type Order struct {
	id int
	items []int
	total float64
	address string
	completed bool
}

	func addOrder(orders map[int]Order, order Order) {
		orders[order.id] = order
		fmt.Println("заказ добавлен:", order.id)
	}

	func main() {
		orders := make(map[int]Order)

		addOrder(orders, Order{
			id: 1,
			items: []int{101,102,103},
			total: 2550,
			address: "Пушкина 1",
			completed: false,
		})

		addOrder(orders, Order{
			id: 2,
			items: []int{201,202},
			total: 3000,
			address: "Пушкина 232",
			completed: true,
		})

		fmt.Println("Все заказы:")
		for id, order := range orders {
			fmt.Println(id, order)
		}
 
	}
