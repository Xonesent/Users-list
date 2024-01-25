package repository

import (
	"context"
	r_api "users-list/internal/repository/api"
	"users-list/server"

	"github.com/jmoiron/sqlx"
)

type People interface {
	Get_Person()
	Delete_Person()
	Patch_Person()
	Post_Person(ctx context.Context, data *server.Person_structure) (int, error)
}

type Repository struct {
	People
}

func New_Repository(db *sqlx.DB) *Repository {
	return &Repository{
		People: r_api.New_People_repository(db),
	}
}
