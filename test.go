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
	// вычитание строк
	if operator == "-" {
		return firstValue
	}

	return ""

}

// 1. Проверяем циклом является ли подстрока частью строки1. Проходимся циклом по каждой букве 1 строки до конца длины подстроки. И сравниваем буквы, пока не найдем совпадение.
// 2. Если нашли подстроку, то создаем новую строку без нее. Иначе возвращаем просто строку.
// 3. Условие firstStrLen-secondStrLen гарантирует, что цикл j не выйдет за пределы. Когда мы ищем подстроку, то максимальный i должен быть началом secondStr, иначе выйдет за рамки.

func stringSubtraction(firstStr, secondStr string) string {
	firstStrLen := len(firstStr)
	secondStrLen := len(secondStr)

	isMatch := true

	for i := 0; i < firstStrLen-secondStrLen; i++ {
		for j := 0; j < secondStrLen; j++ {
			if firstStr[i+j] != secondStr[j] {
				isMatch = false
				fmt.Println(string(firstStr[j]), string(secondStr[j]))
				fmt.Println(isMatch)
			}
		}
	}

	return ""
}

func main() {
	//fmt.Println(getValuesFromInput())
	//fmt.Println(calculate(getValuesFromInput()))
	//fmt.Println(stringSubtraction("world", "wo"))
	fmt.Println(stringSubtraction("Hi World!", "World!"))
}
