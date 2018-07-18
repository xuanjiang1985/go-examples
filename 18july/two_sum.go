package main

import (
	"fmt"
)

func main() {
	s1 := []int{4, 10, 5, 2, 10, 11}
	target := 13
	fmt.Println(twoSum(s1, target))
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i := range nums {
		if pos, ok := mp[target-nums[i]]; ok {
			return []int{pos, i}
		}
		fmt.Println(mp)
		mp[nums[i]] = i
	}
	return []int{}
}
