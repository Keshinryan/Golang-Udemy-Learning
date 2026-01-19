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


func TestPanic(t *testing.T) {
	router:= httprouter.New()
	router.PanicHandler=func(w http.ResponseWriter, r *http.Request, Error interface{}) {
		fmt.Fprint(w, "Panic: ", Error)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		panic("Error")
	})

	request:= httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response:= recorder.Result()
	body,_:= io.ReadAll(response.Body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Panic: Error", string(body))
}

