package test

import (
	"fmt"
	"golang_dependency_injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T){
	simpleService,err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
	//fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceError(t *testing.T){
	simpleService,err := simple.InitializedService(true)
	assert.NotNil(t,err)
	assert.Nil(t,simpleService)
}

func TestSimpleServiceSuccess(t *testing.T){
	simpleService,err := simple.InitializedService(false)
	assert.Nil(t,err)
	assert.NotNil(t,simpleService)
}

