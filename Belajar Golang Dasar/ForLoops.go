package main

import "fmt"

func main() {
	
	//Post Statement
	counter := 1
	for counter <= 10 {
		println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")

	//Init Statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	
	names:=[]string{"Eko","Kurniawan","Khannedy"}
	//For Array / Slice
	for i:=0; i < len(names); i++ {
		fmt.Println("Index", i, "=", names[i])
	}

	//For Range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}