package helper

import (
	"golang_dependency_injection/model/domain"
	"golang_dependency_injection/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse{
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category)[]web.CategoryResponse{
	var categoriesResponses []web.CategoryResponse
	for _,category :=range categories{
		categoriesResponses=append(categoriesResponses, ToCategoryResponse(category))
	}
	return categoriesResponses
}