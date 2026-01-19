package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		Firstname: "jason",
		Lastname:  "patrick",
		Age:       23,
		Hobbies:   []string{"Game", "Reading", "Watching"},
	}
	bytesss, _ := json.Marshal(customer)
	fmt.Println(string(bytesss))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Firstname": "jason", "Lastname":"patrick", "Age": 23, "Hobbies":["Game","Reading","Watching"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname: "jason",
		Lastname:  "patrick",
		Age:       23,
		Hobbies:   []string{"Game", "Reading", "Watching"},
		Address: []Address{
			{
				Street:     "Jl ...",
				Country:    "Indonesia",
				PostalCode: "90870",
			},
			{
				Street:     "Jl ...",
				Country:    "Indonesia",
				PostalCode: "90870",
			},
		},
	}

	bytesss, _ := json.Marshal(customer)
	fmt.Println(string(bytesss))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"jason","Lastname":"patrick","Age":23,"Hobbies":["Game","Reading","Watching"],"Address":[{"Street":"Jl ...","Country":"Indonesia","PostalCode":"90870"},{"Street":"Jl ...","Country":"Indonesia","PostalCode":"90870"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Address)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl ...","Country":"Indonesia","PostalCode":"90870"},{"Street":"Jl ...","Country":"Indonesia","PostalCode":"90870"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	Addresses := []Address{
		{
			Street:     "Jl ...",
			Country:    "Indonesia",
			PostalCode: "90870",
		},
		{
			Street:     "Jl ...",
			Country:    "Indonesia",
			PostalCode: "90870",
		},
	}

	bytesss, _ := json.Marshal(Addresses)
	fmt.Println(string(bytesss))
}
