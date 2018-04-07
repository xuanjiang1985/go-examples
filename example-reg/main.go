package main

import (
	"fmt"
	"regexp"
)

func main(){
	age := " 32岁 agd"
	reg := regexp.MustCompile("[0-9]")
	ageSlice := make([]string, 0, 2)
	ageSlice = reg.FindAllString(age, -1)
	fmt.Println(ageSlice)
	// reg := regexp.MustCompilePOSIX(`[[:word:]].+ `)
	// fmt.Printf("%q\n", reg.FindString("Hello World!"))
}