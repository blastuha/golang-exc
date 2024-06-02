package main

import (
	"fmt"
	"strings"
)

//Вход: "<div><p></p></div>"
//Выход: true
//
//Вход: "<div><p></div></p>"
//Выход: false
//
//Вход: "<div><span></span><p></p></div>"
//Выход: true
//
//Вход: "<div><p></span></p></div>"
//Выход: false

func rightClosedTags(str string) bool {
	stack := []string{}

	for i := 0; i < len(str); i++ {
		// поиск начала тега
		if string(str[i]) == "<" {
			tag := ""
			// поиск закрытия тега

			for j := i; j < len(str); j++ {
				if string(str[j]) == ">" {
					tag = str[i : j+1]

					if !strings.HasPrefix(tag, "</") {
						stack = append(stack, tag)
						break
					} else {
						if len(stack) == 0 {
							return false
						}
						openTag := stack[len(stack)-1][1:len(stack[len(stack)-1])]
						closeTag := tag[2:len(tag)]

						if openTag == closeTag {
							//fmt.Println("open, close", openTag, closeTag)
							stack = stack[0 : len(stack)-1]
							break
						} else {
							return false
						}

					}

				}

			}
			//fmt.Println("stack", stack)

		}
	}
	//fmt.Println(stack)

	return len(stack) == 0
}

func main() {
	fmt.Println(rightClosedTags("<div><p></span></p></div>"))       // false /////
	fmt.Println(rightClosedTags("<div><p></p></div>"))              // true
	fmt.Println(rightClosedTags("<div><p></div></p>"))              // false
	fmt.Println(rightClosedTags("<div><span></span><p></p></div>")) // true
	fmt.Println(rightClosedTags("<div><p></p></div></div>"))        // false
	fmt.Println(rightClosedTags("<p></p>"))                         // true
	fmt.Println(rightClosedTags("<div><div><p></p></div></div>"))   // true
	fmt.Println(rightClosedTags("<div><div><p></div></p></div>"))   // false
}
