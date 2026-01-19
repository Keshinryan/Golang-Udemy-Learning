package main

import "fmt"

type Filter func(string) string
//func sayHelloWithFilter(name string, filter func(string) string) {

//Alternative way to declare function type
func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}