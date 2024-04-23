package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	primer := strings.Split(text, " ")

	if primer[0] == "0" || primer[2] == "0" {
		panic("Диапазон чисел должен быть от 1 до 10")
	}

	if len(primer) < 3 {
		panic("Строка не является математической операцией")
	}

	if len(primer) > 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	rim := false
	res := 0
	znak := primer[1]
	a, _ := strconv.Atoi(primer[0])
	b, _ := strconv.Atoi(primer[2])

	if a == 0 && b == 0 {
		a = romanToInt(primer[0])
		b = romanToInt(primer[2])
		rim = true
	}

	if a == 0 || b == 0 {
		panic("Нельзя использовать одновременно разные системы счисления")
	}

	if rim == true && b > a {
		panic("Результат не может быть отрицательным, так как в римской системе нет отрицательных чисел")
	}

	if a > 10 || a < 1 || b > 10 || b < 1 {
		panic("Диапазон чисел должен быть от 1 до 10")
	}

	switch znak {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		panic("Доступны только операции сложения, вычитания, умножения и деления")
	}

	if rim == false {
		fmt.Println(res)
	}

	if rim == true {
		roman := intToRoman(res)
		fmt.Println(roman)
	}
}

func romanToInt(s string) int {
	roman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	total := roman[s[n-1]]

	for i := n - 2; i >= 0; i-- {
		if roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else {
			total += roman[s[i]]
		}
	}

	return total
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i, val := range values {
		for num >= val {
			result += romans[i]
			num -= val
		}
	}

	return result
}
