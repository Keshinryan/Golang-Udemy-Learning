package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}

	var address2 *Address = &address1 //copy value

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2= &Address{"Jakarta","DKI Jakarta","Indonesia"} hanya mengubah value pointer
	*address2= Address{"Jakarta","DKI Jakarta","Indonesia"} // mengubah value reference dan pointer
	fmt.Println(address1)
	fmt.Println(address2)
}