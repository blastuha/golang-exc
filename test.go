package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var operatorSlice = []string{"+", "-", "*", "/"}

// getValuesFromInput: возвращает 2 значения + оператор
func getValuesFromInput() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")

	line, _ := reader.ReadString('\n')
	noSpaceLine := strings.ReplaceAll(line, " ", "")
	clearLine := strings.TrimSpace(noSpaceLine)

	firstValue := ""
	secondValue := ""
	operator := ""

	for index, value := range clearLine {
		if slices.Contains(operatorSlice, string(value)) {
			firstValue = clearLine[:index]
			secondValue = clearLine[index+1:]
			operator = string(clearLine[index])
		}
	}

	return firstValue, secondValue, operator
}

func calculate(firstValue, secondValue, operator string) string {
	// проверка firstValue на число
	_, err1 := strconv.Atoi(firstValue)
	if err1 == nil {
		panic("First value is a number")
	}
	// перевод secondValue в число и проверка числа на  > 10
	secondNumber, err2 := strconv.Atoi(secondValue)
	if err2 != nil {
		fmt.Println("Второе значение - это строка", err2)
	}
	if secondNumber > 10 {
		panic("Second value as number > 10")
	}
	// проверка values на длину строки > 10
	if len(firstValue) > 10 || len(secondValue) > 10 {
		panic("Длина одного из значений > 10")
	}

	// сложение со строкой или числом
	if operator == "+" {
		return firstValue + secondValue
	}
	// сложение с числом
	if operator == "-" {
		return firstValue + secondValue
	}

	return ""

}

func main() {
	//fmt.Println(getValuesFromInput())
	fmt.Println(calculate(getValuesFromInput()))
}
