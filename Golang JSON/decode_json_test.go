package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T){
	jsonString:=`{"Firstname": "jason", "Lastname":"patrick", "Age": 23}`
	jsonBytes:=[]byte(jsonString)
	
	customer:=&Customer{}
	err:=json.Unmarshal(jsonBytes,customer)
	if err != nil{
		panic(err)
	}
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)

}