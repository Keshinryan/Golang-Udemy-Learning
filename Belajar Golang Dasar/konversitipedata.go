package main

import "fmt"
func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	
	fmt.Println("nilai int32 =", nilai32)
	fmt.Println("nilai int64 =", nilai64)
	fmt.Println("nilai int16 =", nilai16)

	var name ="Jason Patrick"
	var e byte = name[0]
	fmt.Println("name =", name)
	fmt.Println("e =", string(e))
}