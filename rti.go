package main

import "fmt"

//https://leetcode.com/problems/roman-to-integer/
// мое личное решение через stack, но можно решиить намного проще

var romanNumerals2 = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// если число меньше следующего, то из большего вычитается меньшее
// если число болльше следующего или равно, то они складываются

func romanToArab2(str string) int {
	stack := []int{}
	result := 0

	for _, letter := range str {
		romanValue, _ := romanNumerals2[string(letter)]
		stack = append(stack, romanValue)
		//fmt.Println("stack", stack)

		if len(stack) <= 1 {
			result += romanValue
			continue
		}

		if romanValue < stack[len(stack)-2] || romanValue == stack[len(stack)-2] {
			result = result + romanValue
		} else {
			result = result + romanValue - 2*stack[len(stack)-2]

		}

	}

	return result
}

func main() {
	fmt.Println(romanToArab2("IV")) // 4 // -1 + 5
	//fmt.Println(romanToArab2("CXC")) // 190 // 100 - 10 + 100

	fmt.Println(romanToArab2("CXC"))     // 190 // 100 + 10 + 100 - 10 - 10 // 110 + (100 - 10 - 10)
	fmt.Println(romanToArab2("MCM"))     // 1900
	fmt.Println(romanToArab2("MCMXCIV")) // 1994
	fmt.Println(romanToArab2("III"))     // 3
	fmt.Println(romanToArab2("LVIII"))   // 58
}
