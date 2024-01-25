package s_api

import (
	"context"
	"users-list/internal/repository"
	s_help "users-list/internal/service/helpers"
	"users-list/server"
)

type People_service struct {
	people repository.People
}

func New_People_service(people repository.People) *People_service {
	return &People_service{people: people}
}

func (s *People_service) Get_People(ctx context.Context, data *server.Get_structure) ([]server.Patch_structure, error) {
	people, err := s.people.Get_People(ctx, data)
	return people, err
}

func (s *People_service) Delete_Person(ctx context.Context, index int) error {
	err := s.people.Delete_Person(ctx, index)
	return err
}

func (s *People_service) Patch_Person(ctx context.Context, data *server.Patch_structure) error {
	err := s.people.Patch_Person(ctx, data)
	return err
}

func (s *People_service) Post_Person(ctx context.Context, data *server.Post_structure) (int, error) {
	ans, err := s_help.Enriche(ctx, data.Name)
	if err != nil {
		return -1, err
	}

	person := &server.Person_structure{
		Name:        data.Name,
		Surname:     data.Surname,
		Patronymic:  data.Patronymic,
		Age:         ans.Age,
		Gender:      ans.Gender,
		Nationality: ans.Nationality,
	}

	id, err := s.people.Post_Person(ctx, person)
	if err != nil {
		return -1, err
	}

	return id, nil
}
