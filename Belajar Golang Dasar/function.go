package main

import "fmt"

func SayHello() {
	println("Hello")
}

func SayHelloto(firstname string, lastname string) {
	println("Hello", firstname, lastname)
}

// Funtion Return
func getHello(name string) string {
	return "Hello " + name
}

// Function Return Multiple Value
func getFullName() (string, string) {
	return "Jason", "Patrick"
}

func getCompleteName() (firstName, MiddleName, lastName string) {
	firstName = "Eko"
	MiddleName = "Kurniawan"
	lastName = "Khaneddy"

	return firstName, MiddleName, lastName
}

func main() {
	SayHello()
	SayHello()
	SayHello()

	SayHelloto("Eko", "Kurniawan")
	SayHelloto("Budi", "Santoso")

	helloMessage := getHello("Khannedy")
	println(helloMessage)

	firstname, lastname := getFullName()
	println(firstname, " ", lastname)

	firstname, _ = getFullName()
	println(firstname)

	a, b, c := getCompleteName()
	fmt.Println(a,b,c)
}