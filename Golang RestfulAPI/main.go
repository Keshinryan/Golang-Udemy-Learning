package main

import (
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/helper"
	"golang_restful_api/middleware"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	validate:=validator.New()

	db:=app.NewDB()

	categoryRepository:=repository.NewCategoryRepository()
	categoryService:=service.NewCategoryService(categoryRepository, db, validate)
	CategoryController:=controller.NewCategoryController(categoryService)

	router := app.NewRouter(CategoryController)

	server:= http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err:= server.ListenAndServe()
	helper.PanicErr(err)
}