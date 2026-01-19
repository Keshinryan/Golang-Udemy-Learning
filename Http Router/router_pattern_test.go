package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)


func TestRouterPattternNamedParameter(t *testing.T) {
	router:= httprouter.New()
	router.GET("/product/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, Params httprouter.Params){
		id:=Params.ByName("id")
		itemId:=Params.ByName("itemId")
		text:= "Product ID: " + id + ", Item ID: " + itemId
		fmt.Fprint(w, text)
	})

	request:= httptest.NewRequest("GET", "http://localhost:3000/product/1/item/2", nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response:= recorder.Result()
	body,_:= io.ReadAll(response.Body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Product ID: 1, Item ID: 2", string(body))
}

func TestRouterPattternCatchAllParameter(t *testing.T) {
	router:= httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, Params httprouter.Params){
		image:=Params.ByName("image")
		text:= "Image Path: " + image
		fmt.Fprint(w, text)
	})

	request:= httptest.NewRequest("GET", "http://localhost:3000/images/small/image.png", nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response:= recorder.Result()
	body,_:= io.ReadAll(response.Body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Image Path: /small/image.png", string(body))
}

