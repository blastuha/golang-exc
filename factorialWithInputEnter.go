package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Задание придумал gpt по моему запросу. Чтобы был инпут + факториал

func checkString(str string) string {
	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	stringStack := []string{}

	for _, value := range str {
		letter := string(value)
		isLetterNumber := slices.Contains(numbers, letter)

		if isLetterNumber {
			stringStack = append(stringStack, letter)
		}

	}

	return strings.Join(stringStack, "")
}

func getInputString() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите число:")
	scanner.Scan()
	numberString := scanner.Text()
	numberStringOnlyNumbers := checkString(numberString)

	numInt, err := strconv.Atoi(numberStringOnlyNumbers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ваше число:", numInt)
	return numInt
}

// Факториал числа n — это произведение всех натуральных чисел от единицы до n.
// 3! = 1 * 2 * 3

func getFactorial(number int) int {

	if number <= 0 {
		return 1
	}

	return number * getFactorial(number-1)
}

func main() {
	fmt.Println(checkString("1w23qwe"))
	fmt.Println(getFactorial(getInputString()))
}
