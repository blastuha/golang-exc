package main

import (
	"fmt"
	"sort"
)

//["flower","flow","flight"]
// https://leetcode.com/problems/longest-common-prefix/solutions/3539074/go-solution-beats-100-runtime-56-memory/

func commonPrefix(array []string) string {
	firstWord := array[0]

	for i := 1; i < len(array); i++ {
		word := array[i]
		commonLetterCount := 0

		for j := 0; j < len(firstWord) && j < len(word); j++ {
			if firstWord[j] == word[j] {
				fmt.Println("firstWord[j], word[j]", string(firstWord[j]), string(word[j]))
				commonLetterCount++
			} else {
				break
			}
		}
		firstWord = firstWord[0:commonLetterCount]
		fmt.Println("commonLetterCount", commonLetterCount)
	}

	return firstWord
}

func sortArray(array []string) []string {
	sort.Strings(array)
	return array
}

func main() {
	//fmt.Println(commonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(commonPrefix([]string{"flower", "flower", "flower"}))
	fmt.Println(sortArray([]string{"flower", "flow", "flight"}))
}
