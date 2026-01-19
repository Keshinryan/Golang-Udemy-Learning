package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "validation error"}
	}
	if id != "eko" {
		return &notFoundError{Message: "data not found"}
	}
	return nil
}

func main() {
	err := SaveData("budi", nil)

	if err != nil {
		// if ValidationError,ok:= err.(*ValidationError); ok{
		// 	fmt.Println("Validation error:" ,ValidationError.Error())
		// }else if notFoundError,ok:= err.(*notFoundError); ok{
		// 	fmt.Println("Not Found error:" ,notFoundError.Error())
		// }else{
		// 	fmt.Println("Unknown error:", err.Error())
		// }
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("Validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}

}
