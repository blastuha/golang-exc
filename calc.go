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

func isStrRomanNumber(str string) bool {
	for _, letter := range str {
		_, exists := romanNumerals[string(letter)]

		if !exists {
			return false
		}
	}

	return true
}

func romanToArab(str string) int {
	stack := []int{}
	result := 0

	for _, letter := range str {
		romanValue, _ := romanNumerals[string(letter)]
		stack = append(stack, romanValue)

		if len(stack) <= 1 {
			result += romanValue
			continue
		}

		if len(stack) > 1 && romanValue == stack[len(stack)-2] || romanValue < stack[len(stack)-2] {
			result = result + romanValue
		} else {
			result = romanValue - result
		}

	}

	return result
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

	num1 := input[:operatorIndex]
	num2 := input[operatorIndex+1:]
	operator := string(input[operatorIndex])

	if isStrRomanNumber(num1) && isStrRomanNumber(num2) {
		num1 := romanToArab(num1)
		num2 := romanToArab(num2)
		return num1, operator, num2
	}

	num1int, num1Err := strconv.Atoi(num1)
	if num1Err != nil {
		panic(num1Err)
	}
	if num1int > 10 {
		panic("num1 > 10!!")
	}

	num2int, num2Err := strconv.Atoi(num2)
	if num2Err != nil {
		panic(num2Err)
	}
	if num2int > 10 {
		panic("num2 > 10!!")
	}

	return num1int, operator, num2int
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
	// нужно сделать проверку на ввод XXX или IIII
	fmt.Println(calculator())
}
