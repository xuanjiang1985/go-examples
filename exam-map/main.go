package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12008
	personSalary["gang"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
	value, ok := personSalary["gang1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在key")
	}
}
