package main

import "fmt"

// число фибиначи - это последовательность чисел, где каждое число равно сумме двух предыдущих
// функция принимает номер числа в последовательности
// возвращает число фибоначчи
// 0 1 1 2 3 5 8 13...

func fib2(n int) int {
	if n < 2 {
		return n
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {

		result := a + b
		a = b
		b = result

	}

	return b
}

func main() {
	fmt.Println(fib2(5))
}
