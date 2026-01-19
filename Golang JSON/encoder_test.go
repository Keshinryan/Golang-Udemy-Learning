package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T){
	writer, _:= os.Create("sample_output.json")
	encoder:=json.NewEncoder(writer)

	customer:=Customer{
		Firstname: "Jason",
		Lastname: "Patrick",
	}

	encoder.Encode(customer)
}

 