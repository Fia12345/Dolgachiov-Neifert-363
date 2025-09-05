package main

import "fmt"

func main() {
	var osn, ruch, dop float64
	fmt.Println("Основной багаж?")
	fmt.Scan(&osn)

	fmt.Println("Ручная кладь?")
	fmt.Scan(&ruch)

	fmt.Println("Доп. ручная кладь?")
	fmt.Scan(&dop)

	res := osn + ruch + dop

	fmt.Println(res)
}