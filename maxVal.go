package main

import "fmt"

// Задача из GPT
// Для массива [3, 5, 7, 2, 8, -1, 4, 10, 12] максимальный элемент — 12.

func bubleSorting(array []int) []int {

	for i := 0; i < len(array)-1; i++ {

		for j := 0; j < len(array)-1-i; j++ {

			if array[j] > array[j+1] {
				// решение через параллельное присваивание
				array[j+1], array[j] = array[j], array[j+1]
				// решение через временную переменную
				// temp := array[j+1
				//array[j+1] = array[j]
				//array[j] = temp

			}
		}
	}

	return array
}

func findMaxValue(array []int) int {

	return bubleSorting(array)[len(array)-1]
}

func main() {
	fmt.Println(bubleSorting([]int{3, 5, 7, 2, 8, -1, 12, 4, 10}))
	fmt.Println(findMaxValue([]int{3, 5, 7, 2, 8, -1, 12, 4, 10}))
}
