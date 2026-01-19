package middleware

import (
	"golang_dependency_injection/helper"
	"golang_dependency_injection/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	if "RAHASIA" == request.Header.Get("x-API-Key"){
		middleware.Handler.ServeHTTP(writer,request)
	}else{
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unathorized",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}