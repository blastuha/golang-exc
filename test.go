package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var operatorSlice = []string{"+", "-", "*", "/"}

func getValuesFromInput() {
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
			fmt.Println("noSpaceLine[index]", noSpaceLine[index])
		}
	}

	fmt.Println("firstValue, secondValue, operator", firstValue, secondValue, operator)

}

func main() {
	getValuesFromInput()
}
