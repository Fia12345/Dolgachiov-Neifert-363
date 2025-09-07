package main

import (
	"fmt"
	"strconv"
)

const (
	dec = 10
	bin = 2
	hex = 16
)

func convertNumber(numberStr string, fromBase int, toBase int) (string, error) {
	num, err := strconv.ParseInt(numberStr, fromBase, 64)
	if err != nil {
		return "", err
	}
	
	result := strconv.FormatInt(num, toBase)
	return result, nil
}

func main() {
	var decimalNumber int64 = 255
	
	fmt.Printf("Исходное десятичное число: %d\n", decimalNumber)
	binary := strconv.FormatInt(decimalNumber, bin)
	fmt.Printf("Двоичное представление: %s\n", binary)
	hexadecimal := strconv.FormatInt(decimalNumber, hex)
	fmt.Printf("Шестнадцатеричное представление: %s\n", hexadecimal)
	fmt.Println("Пример:")
	result1, _ := convertNumber("42", dec, bin)
	fmt.Printf("42 (10) -> %s (2)\n", result1)
	result2, _ := convertNumber("255", dec, hex)
	fmt.Printf("255 (10) -> %s (16)\n", result2)
	result3, _ := convertNumber("1010", bin, dec)
	fmt.Printf("1010 (2) -> %s (10)\n", result3)
	result4, _ := convertNumber("FF", hex, dec)
	fmt.Printf("FF (16) -> %s (10)\n", result4)
	
	fmt.Println("Введите число для конвертации (или выход):")
	
	for {
		var input string
		fmt.Print("Введите число: ")
		fmt.Scanln(&input)
		
		if input == "выход" {
			break
		}
		
		if result, err := convertNumber(input, dec, bin); err == nil {
			fmt.Printf("%s (10) -> %s (2)\n", input, result)
			continue
		}
		
		if result, err := convertNumber(input, bin, dec); err == nil {
			fmt.Printf("%s (2) -> %s (10)\n", input, result)
			continue
		}
		
		if result, err := convertNumber(input, hex, dec); err == nil {
			fmt.Printf("%s (16) -> %s (10)\n", input, result)
			continue
		}
		
		fmt.Printf("Не удалось распознать число: %s\n", input)
	}
}