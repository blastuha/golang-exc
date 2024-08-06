package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция перевода римских чисел в арабские
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var arabic int
	lastValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[roman[i]]
		if value < lastValue {
			arabic -= value
		} else {
			arabic += value
		}
		lastValue = value
	}
	return arabic, nil
}

// Функция перевода арабских числе в римские
func arabicToRoman(arabic int) (string, error) {
	if arabic < 1 {
		return "", fmt.Errorf("невозможно конвертировать")
	}
	var roman strings.Builder
	numerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, numeral := range numerals {
		for arabic >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return roman.String(), nil
}

// Проверка строки на арабские числа
func isRomanNumber(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !strings.ContainsRune("IVXLCDM", r)
	}) == -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 5 + 3 или V + III):")
	scanner.Scan()
	inputString := scanner.Text()

	valuesArray := strings.Fields(inputString)
	if len(valuesArray) != 3 {
		panic("Неверный формат ввода")
		return
	}

	inputValue1, err1 := strconv.Atoi(valuesArray[0])
	inputValue2, err2 := strconv.Atoi(valuesArray[2])
	inputOperator := valuesArray[1]

	isRoman1 := isRomanNumber(valuesArray[0])
	isRoman2 := isRomanNumber(valuesArray[2])

	if isRoman1 != isRoman2 {
		panic("Числа должны быть либо оба римские, либо оба арабские")
	}

	var result int
	if isRoman1 && isRoman2 {
		inputValue1, err1 = romanToArabic(valuesArray[0])
		inputValue2, err2 = romanToArabic(valuesArray[2])
	}

	if err1 != nil || err2 != nil || inputValue1 < 1 || inputValue1 > 10 || inputValue2 < 1 || inputValue2 > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10!")
		return
	}

	switch inputOperator {
	case "+":
		result = inputValue1 + inputValue2
	case "-":
		result = inputValue1 - inputValue2
	case "*":
		result = inputValue1 * inputValue2
	case "/":
		if inputValue2 == 0 {
			panic("На ноль делить нельзя!")
			return
		} else {
			result = inputValue1 / inputValue2
		}
	default:
		panic("Неверный оператор")
		return
	}

	// Проверка, если ввод римскими то и результат римскими
	if isRoman1 && isRoman2 {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			panic("Ошибка вычисления.")
			return
		}
		fmt.Printf("Ваш результат: %s\n", romanResult)
	} else {
		fmt.Printf("Ваш результат: %d\n", result)
	}
}
