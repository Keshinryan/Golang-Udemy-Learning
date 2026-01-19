package main

import "fmt"
type Address struct {
	City, Province, Country string
}

func ChangeCountrytoIndonesia(address *Address){
	address.Country="Indonesia"

}
func main(){
	var address Address= Address{}
	ChangeCountrytoIndonesia(&address)
	fmt.Println(address)
}