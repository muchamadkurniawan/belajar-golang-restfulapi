package repository

import (
	"belajar-golang-restfulapi/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Sava(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
