package service

import (
	"context"
	"database/sql"
	"golang_restful_api/exception"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
	"golang_restful_api/model/web"
	"golang_restful_api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository, 
		DB: DB, 
		Validate: validate,
	}
}

func (Service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := Service.Validate.Struct(request)
	helper.PanicErr(err)

	tx, err := Service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)
	category := domain.Category{
		Name: request.Name,
	}
	category = Service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
func (Service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := Service.Validate.Struct(request)
	helper.PanicErr(err)

	tx, err := Service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := Service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil{
		panic( exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = Service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)

}
func (Service *CategoryServiceImpl) Delete(ctx context.Context, CategoryId int) {
	tx, err := Service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)
	category, err := Service.CategoryRepository.FindById(ctx, tx, CategoryId)
	if err != nil{
		panic( exception.NewNotFoundError(err.Error()))
	}

	Service.CategoryRepository.Delete(ctx, tx, category)
}
func (Service *CategoryServiceImpl) FindById(ctx context.Context, CategoryId int) web.CategoryResponse {
	tx, err := Service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := Service.CategoryRepository.FindById(ctx, tx, CategoryId)
	if err != nil{
		panic( exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}
func (Service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := Service.DB.Begin()
	helper.PanicErr(err)
	defer helper.CommitOrRollback(tx)

	categories := Service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
