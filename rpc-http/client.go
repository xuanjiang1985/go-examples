package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "ip:port")
		os.Exit(1)
	}

	addr := os.Args[1]

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialhttp: ", err)
	}

	var reply *string

	clientMsg := "I am client, I sent you a message. !!!!!!"
	err = client.Call("MyRPC.HelloRPC", clientMsg, &reply)
	if err != nil {
		log.Fatal("call hellorpc: ", err)
	}
	fmt.Println(*reply)
}
