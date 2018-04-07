package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}
func main() {
	u := User{1, "zhougang", 32}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields", v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)

	}
}


I am a full stack web developer. I use php, golang, mysql, pgsql, redis, mongodb, elasticsearch, sphinx , nginx, docker and so on for the backend of application, and also using goRPC for micro services.
I can design UI photos by myself using PS and I write front-end pages using jquery, angularjs, vuejs, html, css, javascripte and bootstrap.