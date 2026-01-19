package golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T){
	var result string= Length[string]("Hello World")
	assert.Equal(t,result,"Hello World")
	var resultNumber int = Length[int](100)
	assert.Equal(t,resultNumber,100)
}

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2){
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T){
	MultipleParameter[string,int]("eko",100)
	MultipleParameter[int,string](100,"Eko")
}

func IsSame[T comparable](value1, value2 T)bool{
	if value1 == value2{
		return true
	}else{
		return false
	}
}


func TestIsSame(t *testing.T){
	assert.True(t, IsSame[string]("Eko","Eko"))
	assert.True(t, IsSame[int](100,100))
	assert.True(t, IsSame[bool](true,true))
}