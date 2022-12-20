package service

import (
	"belajar-golang-restfulapi/helper"
	"belajar-golang-restfulapi/model/domain"
	"belajar-golang-restfulapi/model/web"
	"belajar-golang-restfulapi/repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicifError(err)

	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Sava(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicifError(err)

	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicifError(err)

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicifError(err)

	category = service.CategoryRepository.Update(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicifError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	fmt.Println("ini service")
	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
