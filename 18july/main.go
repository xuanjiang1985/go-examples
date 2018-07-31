package main

import (
	"18july/link"
	"fmt"
)

func main() {
	var head = link.Node{}
	fmt.Println("head:", head)
	first := link.Node{}
	first.Data = 1
	fmt.Println("first:", first)
	head.Next = &first
	first.Next = &link.Node{}
	//link.Insert(&head, &first, 1)
	fmt.Println("length:", link.Length(&head))
	newNode := link.Node{}
	newNode.Data = 3
	link.Insert(&head, &newNode, 1)
	fmt.Println(head)
	fmt.Println("length:", link.Length(&head))
	fmt.Println(link.Get(&head, 4))
}
