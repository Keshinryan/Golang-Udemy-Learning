package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Jason")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Print(data["name"])
	}
}

