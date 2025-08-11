package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 1
	for {
		if i == len(nums) || len(nums) == 1 {
			return i
		}
		if nums[i-1] == nums[i] {
			fmt.Println(nums)
			nums = append(nums[:i-1], nums[i:]...)
			fmt.Println(nums)
			i = 0
		}
		i++
	}
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1}))
}
