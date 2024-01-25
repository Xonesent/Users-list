package repository

import (
	r_api "users-list/internal/repository/api"

	"github.com/jmoiron/sqlx"
)

type People interface {
	Get_Person()
	Delete_Person()
	Patch_Person()
	Post_Person()
}

type Repository struct {
	People
}

func New_Repository(db *sqlx.DB) *Repository {
	return &Repository{
		People: r_api.New_People_repository(db),
	}
}
