package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Необходимо написать функцию, которая принимает два числа в качестве аргументов и возвращает строку с информацией о том, какое число больше - первое, второе или они равны.
//Затем в main считать из консоли два числа от пользователя и вызвать функцию с этими значениями и вывести результат в консоль.

func getNumbersFromInput() (int, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ReplaceAll(input, " ", ""))

	num1Str := string(input[0])
	num2Str := string(input[1])

	fmt.Println(input)

	num1Int, err1 := strconv.Atoi(num1Str)
	num2Int, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		return -1, -1
	}

	return num1Int, num2Int
}

func randomFunc(num1, num2 int) string {

	if num1 == num2 {
		return "они равны"
	}

	if num1 > num2 {
		return strconv.Itoa(num1)
	} else {
		return strconv.Itoa(num2)
	}
}

func main() {
	//getNumbersFromInput()
	//fmt.Println(getNumbersFromInput())
	fmt.Println(randomFunc(getNumbersFromInput()))
}
