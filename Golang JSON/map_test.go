package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Kumo desu ga nanika Vol 4 Light Novel","imageurl":"http://server/product/image/092131.png","price":110000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["imageurl"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T){
	product:=map[string]interface{}{
		"id":"P001",
		"name":"Kumo desu ga nanika Vol 4 Light Novel",
		"imageurl":"http://server/product/image/092131.png",
		"price":110000,
	}

	bytess,_:=json.Marshal(product)
	fmt.Println(string(bytess))
}