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
		a, err = romeToArab(strings.ToUpper(input1))
		if err != nil {
			fmt.Println("ошибка:", err)
			return
		}
	}

	b, err := strconv.Atoi(input2)
	if err != nil {
		b, err = romeToArab(strings.ToUpper(input2))
		if err != nil {
			fmt.Println("ошибка:", err)
			return
		}
	}

	result, err := calc(a, b, action)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	if isRome(input1) && isRome(input2) {
		fmt.Println("Результат:", arabToRome(result))
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

func calc(a, b int, action string) (int, error) {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		switch action {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "*":
			return a * b, nil
		case "/":
			return a / b, nil
		default:
			return 0, errors.New("неизвестный оператор")
		}
	} else {
		return 0, errors.New("неверный диапозон")
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
