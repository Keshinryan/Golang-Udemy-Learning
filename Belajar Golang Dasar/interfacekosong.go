package main

import "fmt"

func Ups() interface{} {
	return "Ups"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)

}