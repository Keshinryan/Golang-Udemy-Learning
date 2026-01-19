package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	var head *list.Element=data.Front()
	fmt.Println(head.Value)

	next:=head.Next()
	fmt.Println(next.Value)

	next=head.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		println(e.Value.(string))
	}
}