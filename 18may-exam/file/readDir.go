package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileInfos, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}

	for i, fileinfo := range fileInfos {
		fmt.Println(i, fileinfo.Name())
	}

	saveData := []byte("zhou gang is 33")
	err = ioutil.WriteFile("test5.txt", saveData, 0666)

	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadFile("test5.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", b)
	}

	work, _ := ioutil.TempDir("", "go-build")
	fmt.Println("临时目录：", work)
}
