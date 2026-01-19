package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")	
	firstName := "Eko"
	lastName := "Khannedy"
	fmt.Println("Nama saya adalah '", firstName, lastName,"'")
	fmt.Printf("Nama saya adalah '%s %s'\n", firstName, lastName)
	fmt.Printf("Nama saya adalah '%[2]s %[1]s'\n", firstName, lastName)
	fmt.Printf("Nilai 1 + 1 = %d\n", 1+1)
	fmt.Printf("Nilai float dari 22/7 = %.2f\n", 22.0/7.0)
}