package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Redirect")
}

func redirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusMovedPermanently)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://www.google.com", http.StatusMovedPermanently)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", redirectTo)
	mux.HandleFunc("/redirect-from", redirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}