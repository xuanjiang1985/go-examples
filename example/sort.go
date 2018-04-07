package main

import (
	"fmt"
	"sort"
)

var (
	data       []int
	input_Data int /* 声明实际变量 */
)

//输入5个整数，请把这5个数分别由小到大和由小到大输出。

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("输入第%d:", 1)
		fmt.Scan(input_Data)
		data = append(data, input_Data)

	}
	sort.Ints(data)
	fmt.Println("小到大输出：", data)

	fmt.Println("大到小输出：", sort.Reverse(sort.IntSlice(data)))
}
