package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Jason Patrick","Jason"))
	fmt.Println(strings.Split("Jason Patrick"," "))
	fmt.Println(strings.ToLower("Jason Patrick"))
	fmt.Println(strings.ToUpper("Jason Patrick"))
	fmt.Println(strings.Trim("       Jason Patrick      "," "))
	fmt.Println(strings.ReplaceAll("Jason Patrick Jason Patrick","Jason","Eko"))
}