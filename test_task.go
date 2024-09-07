package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input1, action, input2 string

	fmt.Scanf("%s %s %s", &input1, &action, &input2)

	// преобразование входных данных в арабские числа
	a, err := strconv.Atoi(input1)
	if err != nil {
		a, err = romeToArab(input1)
		if err != nil {
			panic("ошибка в первой переменной")
		}
	}

	b, err := strconv.Atoi(input2)
	if err != nil {
		b, err = romeToArab(input2)
		if err != nil {
			panic("ошибка во второй переменной")
		}
	}

	result := calc(a, b, action)

	if isRome(input1) && isRome(input2) {
		if result >= 1 {
			fmt.Println("Результат:", arabToRome(result))
		} else {
			panic("значение не может быть меньше нуля")
		}
	} else if isArab(input1) && isArab(input2) {
		fmt.Println("Результат:", result)
	} else {
		fmt.Println("ошибка")
	}
}
// функция перевода из римских в арабские

func romeToArab(s string) (int, error) {
	rome := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	value, exists := rome[s]
	if exists {
		return value, nil
	} else {
		return 0, fmt.Errorf("неправильное римское число или запись: %s", s)
	}
}

// из арабских в римские

func arabToRome(num int) string {
	val := []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	syms := []string{"C", "XC", "L", "XL", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

	var result strings.Builder

	for i := 0; i < len(val); i++ {
		for num >= val[i] && num > 0 {
			num -= val[i]
			result.WriteString(syms[i])
		}
	}
	return result.String()
}

// функция калькулятора

func calc(a, b int, action string) int {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		switch action {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			panic("неизвестный оператор")
		}
	} else {
		panic("неверный диапозон")
	}
}

// проверяет является ли переменная арабской цифрой
func isArab(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// проверяет является ли переменная римской цифрой

func isRome(s string) bool {
	for _, c := range s {
		if c == 'I' || c == 'V' || c == 'X' || c == 'L' || c == 'C' {
			return true
		}
	}
	return false
}
