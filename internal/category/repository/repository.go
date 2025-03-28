package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/aidosgal/mentor/internal/category/data"
)

type Repository interface {
	List(ctx context.Context) ([]*data.Category, error)
}

type repository struct {
	log *slog.Logger
	db  *sql.DB
}

func NewRepository(log *slog.Logger, db *sql.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) List(ctx context.Context) ([]*data.Category, error) {
	query := `
        SELECT id, name 
        FROM categories ORDER BY lft
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*data.Category
	for rows.Next() {
		category := &data.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
