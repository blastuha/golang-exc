package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/description/

// 1. Если открывающая скобка, то мы кладем в стек
// 2. Если закрывающая, то мы проверяем совпадает ли последняя скобка из стека с закрывающей.

func isValid(str string) bool {
	stack := []string{}

	pairs := make(map[string]string)
	pairs["}"] = "{"
	pairs["]"] = "["
	pairs[")"] = "("

	for _, bracket := range str {
		switch string(bracket) {
		case "{", "[", "(":
			fmt.Println("bracket", string(bracket))
			stack = append(stack, string(bracket))
			fmt.Println("stack", stack)
		case "}", "]", ")":
			if len(stack) == 0 || stack[len(stack)-1] != pairs[string(bracket)] {
				fmt.Println("qwe", stack[len(stack)-1], pairs[string(bracket)], string(bracket), stack[len(stack)-1] != pairs[string(bracket)])
				stack.
				return false
			}
		}

	}

	return true
}

func main() {
	fmt.Println(isValid("{[]}"))
	//fmt.Println(isValid("()[]{}}"))
	//fmt.Println(isValid("(]"))
	//fmt.Println(isValid("{{{}}}"))
	//fmt.Println(isValid("][]{}"))
}
