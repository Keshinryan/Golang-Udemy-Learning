package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//duplikasi bukan pointer (Pass by Value)
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// address2 := address1 //copy value

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	//Pointer (Pass by Reference)
	var address1 Address= Address{"Subang", "Jawa Barat", "Indonesia"}

	var address2 *Address = &address1 //copy value

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	


}