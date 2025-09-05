package main
import "fmt"

 func main() {
	bookingDates := []int{9, 10, 11, 12, 13, 14, 25, 26}

	weekdayPrice := 2100
	weekendPrice := 2850

	totalCost := 0

	for _, day := range bookingDates {
		dayOfWeek := (day-1) % 7 
		if dayOfWeek < 4 {
			totalCost += weekdayPrice
		} else {
		totalCost += weekendPrice
		}
	}
	fmt.Println(totalCost)
 }