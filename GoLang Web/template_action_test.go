package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/if.gohtml"))
	data := map[string]interface{}{
		"Title": "Template Action If",
		"Name":  "Alice Johnson",
		"Address": map[string]interface{}{
			"Street": "789 Sample Rd",
		},
	}
	t.ExecuteTemplate(writer, "if.gohtml", data)
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/comparator.gohtml"))
	data := map[string]interface{}{
		"Title": "Template Action Comparator",
		"FinalValue": 70,
	}
	t.ExecuteTemplate(writer, "comparator.gohtml", data)
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateActionComparator(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/range.gohtml"))
	data := map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{"Reading", "Gaming", "Cooking"},
	}
	t.ExecuteTemplate(writer, "range.gohtml", data)
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateActionRange(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/address.gohtml"))
	data := map[string]interface{}{
		"Name": "Bob Williams",
		"Address": map[string]interface{}{
			"Street": "321 Example Ave",
			"City": "New York",
		},
	}
	t.ExecuteTemplate(writer, "address.gohtml", data)
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateActionWith(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}