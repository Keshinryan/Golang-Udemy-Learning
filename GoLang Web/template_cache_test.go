package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed template/*.gohtml
var tmplts embed.FS

var MyTemplates = template.Must(template.ParseFS(tmplts, "template/*.gohtml"))

func TemplateCache(writer http.ResponseWriter, request *http.Request) {
	MyTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Cache")
}

func TestTemplateCache(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateCache(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
