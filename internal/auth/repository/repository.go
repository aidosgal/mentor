package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/aidosgal/mentor/internal/auth/data"
)

type Repository interface {
	Create(ctx context.Context, user *data.UserModel) (int64, error)
	Get(ctx context.Context, chatID string) (*data.UserModel, error)
}

type repository struct {
	db *sql.DB
	log *slog.Logger
}

func NewRepository(db *sql.DB, log *slog.Logger) Repository {
	return &repository{
		db: db,
		log: log,
	}
}

func (r *repository) Create(ctx context.Context, user *data.UserModel) (int64, error) {
	query :=
	`
        INSERT INTO users (
            first_name,
            last_name,
            phone,
            chat_id,
            username,
            role,
            description
        )
	    VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `
	
	var id int64
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.ChatID,
		user.UserName,
		user.Role,
		user.Description,
	).Scan(&id)
	if err != nil { 
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, chatID string) (*data.UserModel, error) {
	query :=
	`
        SELECT id, first_name, last_name, phone, chat_id, username, role, description
        FROM users
        WHERE chat_id = $1
    `
	
	user := &data.UserModel{}
	
	err := r.db.QueryRow(query, chatID).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.ChatID,
		&user.UserName,
		&user.Role,
		&user.Description,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return user, nil
}
