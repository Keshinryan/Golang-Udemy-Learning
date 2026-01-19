package main

import (
	"fmt"
	"errors"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh nol")
	}
	return nilai / pembagi, nil
}

func main() {
	result, err := Pembagian(10, 0)	
	if err != nil {
		fmt.Println("Terjadi error:", err.Error())
	} else {
		fmt.Println("Hasil pembagian:", result)
	}
}