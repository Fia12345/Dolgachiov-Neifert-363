package main
import "fmt"

type Employee struct {
	ID int
	Name string
	Position string
	Salary float64
}

	func calculateSalaryStats(employees []Employee) (float64, float64) {
		total := 0.0
		count := len(employees)

		for _, emp := range employees {
			total += emp.Salary
		}

		average := total / float64(count)

		return total, average
	}

	func main() {
		employees := []Employee{
			{ID: 1, Name: "Вова", Position: "Президент", Salary:9999999},
			{ID: 2, Name: "Константин", Position: "Вьетнам", Salary:10},
			{ID: 3, Name: "Илья", Position: "Чуть ниже Президента", Salary:999999},
			{ID: 4, Name: "Вова2", Position: "Президент2", Salary:9999992},
		}
		fmt.Printf("все сотрудники:")
		for _, emp := range employees {
			fmt.Printf("ID: %d, Имя: %s, Должность: %s, Зарплата: %.2f Руб.\n",emp.ID, emp.Name, emp.Position, emp.Salary)
		}
		total, average := calculateSalaryStats(employees)
		fmt.Printf("Общий фонд ЗП: %.2f Руб. \n", total)
		fmt.Printf("Средняя ЗП: %.2f Руб. \n", average)
	}
