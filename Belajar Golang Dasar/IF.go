package main

import "fmt"

func main() {
	// code for comparison operators will go here

	name := "Joko"

	if name == "Jason" {
		fmt.Println("Hello Jason")
	}else if name == "Joko"{
		fmt.Println("Hi Joko")
	}else{
		fmt.Println("Hi, Boleh Kenalan?")
	}

	//short statement

	if length := len(name); length>5{
		fmt.Println("Nama terlalu Panjang")
	}else{
		fmt.Println("Nama sudah benar")
	}
}