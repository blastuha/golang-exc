package main

import "fmt"

//https://leetcode.com/problems/kth-largest-element-in-an-array/solutions/5057982/two-solution-with-1-sort-method-and-2-buble-sort/?envType=study-plan-v2&envId=top-interview-150

//Input: nums = [2,1,4,3], k = 2
//Output: 3

func findKthLargest2(nums []int, k int) int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			fmt.Println(nums[j], nums[j+1])
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				fmt.Println("swapped ok!", nums)
			}

		}
	}
	fmt.Println("lastNums", nums)
	return nums[len(nums)-k]
}

func main() {
	fmt.Println(findKthLargest2([]int{2, 1, 4, 3}, 2))
}
