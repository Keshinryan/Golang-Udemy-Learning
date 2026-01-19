package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err:=strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}	else{
		fmt.Println("Result:", result)
	}

	resultInt, err:=strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}	else{
		fmt.Println("Result:", resultInt)
	}

	binary := strconv.FormatInt(1000,2)
	fmt.Println("Binary:", binary)

	var stringInt string = strconv.Itoa(1000)
	fmt.Println("String Int:", stringInt)
}