package service

import "users-list/internal/repository"

type People interface {
	Get_Person()
	Delete_Person()
	Patch_Person()
	Post_Person()
}

type Service struct{
	People
}

func New_Service(repos *repository.Repository) *Service {
	return &Service{}
}
