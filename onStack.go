package main

import (
	"fmt"
	"strings"
)

//Вход: s = "abbaca"
//Выход: "ca"
//Объяснение:
//Например, в строке "abbaca" сначала удаляются две соседние 'b', чтобы получить "aaca".
//Затем удаляются две соседние 'a', чтобы получить "ca".

func deleteDubl(str string) string {
	stack := []string{}

	for i := 0; i < len(str); i++ {
		letter := string(str[i])
		if len(stack) == 0 || stack[len(stack)-1] != letter {
			stack = append(stack, letter)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	fmt.Println("stack", stack)
	result := strings.Join(stack, "")
	return result
}

func main() {
	fmt.Println(deleteDubl("abbaca"))
	fmt.Println(deleteDubl("azxxzy"))
}
