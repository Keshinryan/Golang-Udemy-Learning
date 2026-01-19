package main
import "fmt"
func main()  {
	const firstName ="Jason"
	const lastName ="Patrick"

	// firstName="Michael"
	// lastName="Jordan"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const(
		firstName2="Michael"
		lastName2="Jordan"
	)
	fmt.Println(firstName2)
	fmt.Println(lastName2)
}