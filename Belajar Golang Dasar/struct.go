package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//method
func (customer Customer) sayHello(name string){
	fmt.Println("Hello, ",name,". My name is ", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Jason Patrick"
	eko.Address = "Indonesia"
	eko.Age = 22

	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	Joko:= Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(Joko)

	Budi:= Customer{"Budi","Indonesia",30}
	fmt.Println(Budi)

	Budi.sayHello("Agus")
	Joko.sayHello("Agus")
}
