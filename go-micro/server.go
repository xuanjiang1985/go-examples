package main

import (
	"encoding/json"
	"fmt"
	micro "github.com/micro/go-micro"
	"go-micro/greeter"
	"golang.org/x/net/context"
)

type Person struct {
	Id        int    `json:"id"`
	IdCard    string `json:"idcard"`
	Name      string `json:"name"`
	Age       int8   `json:"age"`
	CreatedAt int64  `json:"created_at"`
}

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	msg := []byte(req.Name)
	var data Person
	err := json.Unmarshal(msg, &data)
	if err != nil {
		// 写日志
	}
	fmt.Println(data)
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
