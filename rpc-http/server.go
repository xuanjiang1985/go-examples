package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type MyRPC int

func (r *MyRPC) HelloRPC(clientMsg string, reply *string) error {
	fmt.Println(clientMsg)
	*reply = "I am the server, I got your msg !!!!!!!!!!!"
	return nil
}

func main() {
	r := new(MyRPC)
	rpc.Register(r)
	rpc.HandleHTTP()
	err := http.ListenAndServe("127.0.0.1:1234", nil)
	if err != nil {
		fmt.Println("in main", err.Error())
	}
}
