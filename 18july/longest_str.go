package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLongest("abcd"))
}

func getLongest(str string) int {
	strMap := make(map[string]int, len(str))
	for i := 0; i < len(str); i++ {
		strMap[string(str[i])] = -1
	}
	pre := -1 // 第一次出现位置
	max := 0
	for i := 0; i < len(str); i++ {
		pre = Max(pre, strMap[string(str[i])])
		max = Max(max, i-pre)

		strMap[string(str[i])] = i
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
