package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected", pc.name)
}
func main() {
	a := PhoneConnector{"PhoneName"}
	a.Connect()
	Disconnect(123)
}

func Disconnect(usb interface{}) {
	// if p, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected", p.name)
	// 	return
	// }
	// fmt.Println("Unknown device")
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected", v.name)
	default:
		fmt.Println("Unknown device")
	}
}
