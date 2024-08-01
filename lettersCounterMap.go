package main

import "fmt"

//Написать функцию, которая принимает строку и возвращает количество вхождений каждого символа в этой строке.
//Используйте хеш-таблицу для хранения и подсчета количества вхождений символов.
// 'hello world' {'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1}

func dublLettersCounter(str string) map[string]int {
	var lettersMap = make(map[string]int)

	for _, value := range str {
		_, exist := lettersMap[string(value)]

		if !exist {
			lettersMap[string(value)] = 1
		} else {
			lettersMap[string(value)] = lettersMap[string(value)] + 1
		}
	}

	return lettersMap
}

func main() {
	fmt.Println(dublLettersCounter("boryaa"))
}
