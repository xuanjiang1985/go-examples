package main

import (
	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var addr = flag.String("addr", "localhost:6001", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	// redis
	channel := r.FormValue("channels")
	if channel == "" {
		log.Println("channels:", channel)
		return
	}
	channels := strings.Split(channel, ",")

	rc, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		log.Println(err, " 1")
	}
	defer rc.Close()

	psc := redis.PubSubConn{Conn: rc}

	if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		log.Println(err, " 2")
	}
	msg := make(chan []byte, 4)

	// 服务端推送消息
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Subscription:
				log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case redis.Message: //单个订阅subscribe
				msg <- v.Data
			case error:
				return
			}
		}
	}()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// 客户端传来消息
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			msg <- message
		}
	}()

	for {
		err = c.WriteMessage(1, <-msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
