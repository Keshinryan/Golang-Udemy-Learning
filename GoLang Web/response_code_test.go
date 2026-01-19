package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name:= request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is required")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T){
	request:= httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder:= httptest.NewRecorder()
	ResponseCode(recorder, request)
	response:= recorder.Result()
	body, _:= io.ReadAll(response.Body)
	fmt.Println("Response Code:", response.StatusCode)
	fmt.Println("Response Body:", response.Status)
	fmt.Println("Response Body:", string(body))
}

func TestResponseCodeValid(t *testing.T){
	request:= httptest.NewRequest("GET", "http://localhost:8080?name=John", nil)
	recorder:= httptest.NewRecorder()
	ResponseCode(recorder, request)
	response:= recorder.Result()
	body, _:= io.ReadAll(response.Body)
	fmt.Println("Response Code:", response.StatusCode)
	fmt.Println("Response Body:", response.Status)
	fmt.Println("Response Body:", string(body))
}
