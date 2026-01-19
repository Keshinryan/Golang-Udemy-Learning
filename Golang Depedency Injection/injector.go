//go:build wireinject
// +build wireinject
package main

import (
	"golang_dependency_injection/app"
	"golang_dependency_injection/controller"
	"golang_dependency_injection/middleware"
	"golang_dependency_injection/repository"
	"golang_dependency_injection/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var CategorySet= wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository),new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService),new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server{
	wire.Build(app.NewDB,
		validator.New,
		CategorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler),new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return  nil
}