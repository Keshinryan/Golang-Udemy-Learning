package main

import (
	"Golang/database"
	_ "Golang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
