package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type loggingMiddleware struct {
	handler http.Handler
}

type ErrorMiddleware struct {
	handler http.Handler
}

func (middleware *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// before request
	println("Before Request")
	middleware.handler.ServeHTTP(w, r)
	// after request
	println("After Request")
}

func (middleware *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error occurred:")
			w.WriteHeader( http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()
	middleware.handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T){
	mux:= http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("About Handler executed")
		fmt.Fprint(writer, "About Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Handler executed")
		panic("Panic occurred")
	})
	loggingMiddleware:=&loggingMiddleware{
		handler: mux,
	}
	errorMiddleware:=&ErrorMiddleware{
		handler: loggingMiddleware,
	}
	server:= http.Server{
		Addr: "localhost:8080",
		Handler: errorMiddleware,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

