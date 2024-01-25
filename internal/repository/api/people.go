package r_api

import "github.com/jmoiron/sqlx"

type People_repository struct {
	db *sqlx.DB
}

func New_People_repository(db *sqlx.DB) *People_repository {
	return &People_repository{db: db}
}

func (h *People_repository) Get_Person() {

}

func (h *People_repository) Delete_Person() {

}

func (h *People_repository) Patch_Person() {

}

func (h *People_repository) Post_Person() {

}
