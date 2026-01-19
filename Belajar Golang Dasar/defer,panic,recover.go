package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(){
	defer logging()
	fmt.Println("Run Application")
}

func endApp(){	
	fmt.Println("End App")
    message:=recover()
	fmt.Println("Terjadi Error", message)
}

func runApp (error bool){
	defer endApp()

	if error{
		panic("Ups Error")
	}
	// recover yang salah
	// message:=recover()
	// fmt.Println("Terjadi Error", message)
}



func main() {
	runApplication()
	runApp(true)
}