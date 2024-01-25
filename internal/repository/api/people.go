package r_api

import (
	"context"
	"errors"
	"fmt"
	r_help "users-list/internal/repository/helpers"
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

func (r *People_repository) Delete_Person(ctx context.Context, index int) error {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", Users_Table)
	err := r.db.QueryRow(query, index).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("no id found")
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", Users_Table)
	_, err = r.db.Exec(query, index)
	return err
}

func (r *People_repository) Patch_Person(ctx context.Context, data *server.Patch_structure) error {
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", Users_Table)
	err := r.db.QueryRow(query, data.Id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("no id found")
	}
	
	sql_req, sql_args := r_help.ExtractSQL(data)

	query = fmt.Sprintf("UPDATE %s SET %s", Users_Table, sql_req)
	_, err = r.db.Exec(query, sql_args...)

	return err
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
