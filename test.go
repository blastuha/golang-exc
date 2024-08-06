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

	firstValue := ""
	secondValue := ""
	operator := ""

	for index, value := range noSpaceLine {
		if slices.Contains(operatorSlice, string(value)) {
			firstValue = noSpaceLine[:index]
			secondValue = noSpaceLine[index+1:]
			operator = string(noSpaceLine[index])
		}
	}

	_, err := strconv.Atoi(firstValue)
	if err == nil {
		panic("First value is a number")
	}

	if len(firstValue) > 10 || len(secondValue) > 10 {
		panic("Длина одного из значений > 10")
	}

	return firstValue, secondValue, operator
}

func calculate(firstValue, secondValue, operator string) {
	switch operator {
	case "+":

	}

}

func main() {
	getValuesFromInput()
}
