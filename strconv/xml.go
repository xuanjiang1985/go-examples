package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	p := person{"zhougang", 31}

	var data []byte
	var err error

	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	p2 := new(person)

	if err = xml.Unmarshal(data, &p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2)

}
