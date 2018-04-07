package main

import (
	"flag"
	"fmt"
)

func style() {
	methodPrt := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")

	//解析
	flag.Parse()
	fmt.Println(*methodPrt)
	fmt.Println(*valuePtr)
}

func style2() {
	var method string
	var value int
	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()

	fmt.Println(method)
	fmt.Println(value)
}

func main() {
	style2()
}
