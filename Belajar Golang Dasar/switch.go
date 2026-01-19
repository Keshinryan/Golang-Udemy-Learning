package main

import "fmt"

func main() {
	name := "Eko"
	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Boleh Kenalan?")
	}

	switch length := len(name); length>5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Pas")
	}

	switch {
	case name == "Eko":
		fmt.Println("Hello Eko")
	case name == "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Boleh Kenalan?")
	}
}