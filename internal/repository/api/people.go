package r_api

import (
	"context"
	"fmt"
	"users-list/server"

	"github.com/jmoiron/sqlx"
)

type People_repository struct {
	db *sqlx.DB
}

func New_People_repository(db *sqlx.DB) *People_repository {
	return &People_repository{db: db}
}

func (r *People_repository) Get_Person() {

}

func (r *People_repository) Delete_Person() {

}

func (r *People_repository) Patch_Person() {

}

func (r *People_repository) Post_Person(ctx context.Context, data *server.Person_structure) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age, gender, nationality) values ($1, $2, $3, $4, $5, $6) RETURNING id", Users_Table)

	row := r.db.QueryRow(query, data.Name, data.Surname, data.Patronymic, data.Age, data.Gender, data.Nationality)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
