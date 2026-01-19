package main

import "fmt"

func main() {
	// person:= map[string]string{}
	// person["name"]="Jason"
	// person["address"]="TirtaSiak"

	person := map[string]string{
		"name":    "Jason",
		"address": "TirtaSiak",
	}

	println(person["name"])
	println(person["address"])
	fmt.Println(person)

	book:=make(map[string]string)
	book["title"]="otonari no tenshi"
	book["author"]="Hanekoto"
	book["Ups"]="salah"

	fmt.Println(book)

	delete(book,"Ups")
	fmt.Println(book)
}