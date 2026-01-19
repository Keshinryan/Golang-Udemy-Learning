package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageurl"`
	Price    int32  `json:"price"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Kumo desu ga nanika Vol 4 Light Novel",
		ImageURL: "http://server/product/image/092131.png",
		Price:    110000,
	}

	bytess, _ := json.Marshal(product)
	fmt.Println(string(bytess))
}

func TestJSONTagDecode(t *testing.T) {
	jsonstring:=`{"ID":"P001","NAME":"Kumo desu ga nanika Vol 4 Light Novel","imageurl":"http://server/product/image/092131.png","price":110000}`
	jsonBytes:=[]byte(jsonstring)

	product:=&Product{}

	json.Unmarshal(jsonBytes,product)
	fmt.Println(product)
}
