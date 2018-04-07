// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "localhost:6001", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
} // use default options

var psc redis.PubSubConn

func init() {
	// redis
	rc, err := redis.Dial("tcp", "127.0.0.1:6379")
	// Read timeout on server should be greater than ping period.
	// redis.DialReadTimeout(20*time.Second),
	// redis.DialWriteTimeout(10*time.Second)
	if err != nil {
		fmt.Println(err, " 1")
	}
	defer rc.Close()

	psc = redis.PubSubConn{Conn: rc}

	if err := psc.Subscribe("*"); err != nil {
		fmt.Println(err, " 2")
	}

}

func echo(w http.ResponseWriter, r *http.Request) {
	msg := make(chan []byte, 20)

	go func(ch chan []byte) {
		for {
			switch v := psc.Receive().(type) {
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case redis.Message: //单个订阅subscribe
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
				msg <- v.Data
			case error:
				return
			}
		}
	}(msg)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		//fmt.Println(psc.Receive())
		// log.Printf("recv: %s", message)
		time.Sleep(time.Second)
		message := `{"message":"wang","name":"123"}`
		select {
		case has := <-msg:
			c.WriteMessage(1, has)
		default:
			c.WriteMessage(1, []byte(message))
		}

		// err = c.WriteMessage(1, []byte(message))
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
// }

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
