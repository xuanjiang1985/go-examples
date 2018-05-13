package main

import (
	"encoding/json"
	"fmt"
	micro "github.com/micro/go-micro"
	"go-micro/greeter"
	"golang.org/x/net/context"
	"time"
)

type Person struct {
	Id        int    `json:"id"`
	IdCard    string `json:"idcard"`
	Name      string `json:"name"`
	Age       int8   `json:"age"`
	CreatedAt int64  `json:"created_at"`
}

func main() {
	msg := Person{
		Id:        1,
		IdCard:    "362424198501275911",
		Name:      "周刚",
		Age:       33,
		CreatedAt: time.Now().Unix(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		data = []byte("{code:1004, msg:无数据}")
	}
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter1.client122"))

	// Create new greeter client
	c := greeter.NewGreeterClient("greeter", service.Client())

	// Call the greeter
	rsp, err := c.Hello(context.TODO(), &greeter.HelloRequest{Name: string(data)})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
