package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname string 
	Lastname  string 
	Age       int32  
	Hobbies   []string
	Address	  []Address

}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname: "jason",
		Lastname:  "patrick",
		Age:       23,
	}
	bytess, _:= json.Marshal(customer)
	fmt.Println(string(bytess))
}


