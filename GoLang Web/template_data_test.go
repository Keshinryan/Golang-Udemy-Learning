package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func templateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/name.gohtml"))
	data := map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "John Doe",
		"Address": map[string]interface{}{
			"Street": "456 Another St",
		},
	}
	t.ExecuteTemplate(writer, "name.gohtml", data)
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateDataMap(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type PageData struct {
	Title string
	Name  string
	Address Address
}

func templateDataStruct(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("template/name.gohtml"))
	data := PageData{
		Title: "Template Data Struct",
		Name:  "Jane Smith",
		Address: Address{
			Street: "123 Main Street",
		},
	}
	t.ExecuteTemplate(writer, "name.gohtml", data)
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()	
	templateDataStruct(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
