package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp=regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eno"))
	fmt.Println(regex.MatchString("aio"))

	fmt.Println(regex.FindAllString("eko edo eno aio",10))
}