package main

import (
	"sort"
)

//[1, 2, 10, 8] --> [8, 10]
//[1, 5, 87, 45, 8, 8] --> [45, 87]
//[1, 3, 10, 0]) --> [3, 10]

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	lastIndex := len(ages)
	preLastIndex := len(ages) - 2
	result := ages[preLastIndex:lastIndex]
	return [2]int(result)
}

func main() {
	TwoOldestAges([]int{1, 5, 87, 45, 8, 8})
}
