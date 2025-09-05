package main
import ("fmt"; "errors"; "strings")

func main(){
    testCases := [] struct {
		name string
		age int
		email string
	}{
	{"Ваня", 25, "ivan228@goida.RussianFederation"},
	{"Вова", 740, "ZOVHERO_OF_RUSSIAZOV@kremlin.RussianFederation"},
	{"Дмитрий", 60,"knopkakremlin.Zov" },
    {"", 9, "пшьфшдру@ру.КНР"},
	}
	for _, tc := range testCases {
		fmt.Printf("Проверка: %s, %d, %s\n", tc.name, tc.age, tc.email)
		err := validateUser(tc.name, tc.age, tc.email)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println("Данные правильные")
		}
		fmt.Println("---")
	}
}

	func validateUser(name string, age int, email string) error {
		if name == "" {
			return errors.New("Имя пустое")
		}
		if len(name) >= 50 {
			return errors.New("Имя должно быть меньше 50 символов")
		}

		if age < 18 {
			return errors.New("Возраст от 18 лет")
		}
		if age > 120 {
			return errors.New("Возраст должен быть до 120 лет")
		}

		if !strings.Contains(email, "@") {
			return errors.New("email должен иметь @")
		}
		return nil
	}