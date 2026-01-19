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
type LogMiddleware struct {
	Handler http.Handler
}

func  (lm *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	lm.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router:= httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprint(w, "Middleware")
	})

	middleware:= &LogMiddleware{
		Handler: router,
	}


	request:= httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder:= httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response:= recorder.Result()
	body,_:= io.ReadAll(response.Body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Middleware", string(body))
}

