package main

import (
	"fmt"
)

// Функция пузырьковой сортировки строк по длине и лексикографически
func bubbleStringSorting(arr []string) []string {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			// Сравниваем длины строк
			if len(arr[j]) > len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else if len(arr[j]) == len(arr[j+1]) && arr[j] > arr[j+1] {
				// Если длины строк одинаковые, сравниваем их лексикографически
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleStringSorting([]string{"pie", "banana", "pear", "peach", "plum", "apple"}))
}
