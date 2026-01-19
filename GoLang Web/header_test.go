package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	ContentType:= request.Header.Get("Content-Type")
	fmt.Fprintf(writer, "Content-Type: %s", ContentType)	
}

func TestRequestHeader(t *testing.T) {
	request:= httptest.NewRequest("GET", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder:= httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Golang")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request:= httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder:= httptest.NewRecorder()	
	ResponseHeader(recorder, request)

	response := recorder.Result()
	xPoweredBy := response.Header.Get("X-Powered-By")
	fmt.Println("X-Powered-By:", xPoweredBy)
}