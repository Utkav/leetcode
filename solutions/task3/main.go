package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numsSet := make(map[int]bool)

	for _, val := range nums {
		_, ok := numsSet[val]
		if ok {
			return true
		}
		numsSet[val] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
