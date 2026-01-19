package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jason"
	names[1] = "Patrick"

	println(names[0])
	println(names[1])
	fmt.Println(names)

	var number =[...] int{
		10,
		120,
		1000,
	}

	fmt.Println(number)
	fmt.Println(len(names))
	number[0]=1
	fmt.Println(number[0])

}