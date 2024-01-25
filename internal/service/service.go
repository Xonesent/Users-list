package service

import (
	"context"
	"users-list/internal/repository"
	s_api "users-list/internal/service/api"
	"users-list/server"
)

type People interface {
	Get_People(ctx context.Context, data *server.Get_structure) ([]server.Patch_structure, error)
	Delete_Person(ctx context.Context, index int) error
	Patch_Person(ctx context.Context, data *server.Patch_structure) error
	Post_Person(ctx context.Context, data *server.Post_structure) (int, error)
}

type Service struct {
	People
}

func New_Service(repos *repository.Repository) *Service {
	return &Service{
		People: s_api.New_People_service(repos.People),
	}
}
