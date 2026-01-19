package main

import "fmt"

func factorialloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int{
	if value==1{
		return 1
	}else{
		return value* factorialRecursive(value-1)
	}
}
func main() {
	fmt.Println(factorialloop(10))
	fmt.Println(factorialRecursive(10))
}