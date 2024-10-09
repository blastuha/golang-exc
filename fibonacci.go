package main

import "fmt"

// число фибиначи - это последовательность чисел, где каждое число равно сумме двух предыдущих
// функция принимает номер числа в последовательности
// возвращает число фибоначчи
// Функция будет вызывать сама себя, пока не встретится базовый случай. Тогда она передаст значение по цепочке вверх и рекурсия закончится.
// 0 1 1 2 3 5 8 13...

var cache = make(map[int]int)

func fib(n int) int {

	if n < 2 {
		//cache[n] = n
		return n
	}

	value, ok := cache[n]

	if ok {
		return value
	}

	result := fib(n-1) + fib(n-2)

	cache[n] = result

	return result
}

func main() {
	//fmt.Println(fib(3))
	//fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(cache)
}
