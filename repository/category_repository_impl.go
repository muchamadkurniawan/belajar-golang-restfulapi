package repository

import (
	"belajar-golang-restfulapi/helper"
	"belajar-golang-restfulapi/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Sava(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category (id) values (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicifError(err)

	id, err := result.LastInsertId()
	helper.PanicifError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicifError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicifError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, CategoryId)
	helper.PanicifError(err)
	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(CategoryId, &category.Name)
		helper.PanicifError(err)
		return category, nil
	} else {
		return category, errors.New("Category is not Found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	fmt.Println("ini repository ")
	sql := "select id, name from category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicifError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(category.Id, &category.Id)
		helper.PanicifError(err)
		categories = append(categories, category)
	}
	return categories
}
