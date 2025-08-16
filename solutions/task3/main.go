package main

import "fmt"

func containsDuplicate(nums []int) bool {
	signs_count := make(map[int]int)

	for _, val := range nums {
		_, ok := signs_count[val]
		if ok {
			return true
		} else {
			signs_count[val] = 1
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
