package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string	`required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var field reflect.StructField = valueType.Field(i)
		fmt.Println(field.Name,"with type",field.Type)
		fmt.Println(field.Tag.Get("required"))
		fmt.Println(field.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	t:= reflect.TypeOf(value)
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)
		if f.Tag.Get("required")=="true"{
			data:=reflect.ValueOf(value).Field(i).Interface()
			result=data!=""
			if result==false{
				return result
			}
		}
	}
	return true
}


func main() {
	readField(Sample{"Eko"})
	readField(Person{"Budi","",""})

	person:=Person{"Joko","","tes"}
	fmt.Println(IsValid(person))
}