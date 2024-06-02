package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = [4]string{"+", "-", "*", "/"}
var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10}

func isValidOperator(operator string) bool {
	for _, v := range operators {
		if operator == v {
			return true
		}
	}

	return false
}

func getExpressionValues() (int, string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ReplaceAll(input, " ", ""))

	fmt.Println(input)

	operatorIndex := -1

	for i, v := range input {
		if isValidOperator(string(v)) {
			operatorIndex = i
			break
		}
	}

	if operatorIndex == -1 {
		panic("Оператор не найден")
	}

	num1, num1Err := strconv.Atoi(input[:operatorIndex])
	if num1Err != nil {
		panic(num1Err)
	}
	if num1 > 10 {
		panic("num1 > 10!!")
	}

	num2, num2Err := strconv.Atoi(input[operatorIndex:])
	if num2Err != nil {
		panic(num2Err)
	}
	if num2 > 10 {
		panic("num2 > 10!!")
	}

	operator := string(input[operatorIndex])

	return num1, operator, num2
}

func calculator() int {
	num1, operator, num2 := getExpressionValues()
	fmt.Println(num1, operator, num2)

	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		return num1 / num2
	case "*":
		return num1 * num2
	default:
		panic("Оператор не найден")
	}
}

func main() {
	fmt.Println(calculator())
}
