package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	path := "test.txt"
	content := "hellow world 竹 宣讲🏆"
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile err is :", err)
		return
	}

	fmt.Println([]byte(content))
}
