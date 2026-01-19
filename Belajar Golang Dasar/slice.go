package main

import "fmt"

func main() {
	// code for comparison operators will go here
	name := [...]string{"Jason", "Bill", "Seun", "Himmel", "Yasil", "Booyah"}

	slices1 := name[1:3]
	fmt.Println(slices1)

	slices2 := name[:3]
	fmt.Println(slices2)

	slices3 := name[3:]
	fmt.Println(slices3)

	slices4 := name[:]
	fmt.Println(slices4)

	fmt.Println(len(slices1))
	fmt.Println(cap(slices2))
	slices4[0]="Patrick"
	slices4[1]="Budi"
	fmt.Println(slices4)

	slices4=append(slices4,"Aji")

	fmt.Println(slices4)

	newSlice := make([]string, 2,5)
	newSlice[0]= "Jason"
	newSlice[1]= "Jason"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice2:= append(newSlice,"Jason")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0]="Budi"
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := name[:]
	toSlice := make([]string, len(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	iniArray :=[...]int{1,2,3,4,5}
	iniSlice :=[]int{6,7,8,9,10}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}