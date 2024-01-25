package s_api

import (
	"context"
	"fmt"
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

func (h *People_service) Get_Person() {

}

func (h *People_service) Delete_Person() {

}

func (h *People_service) Patch_Person() {

}

func (h *People_service) Post_Person(ctx context.Context, data *server.Post_structure) (int, error) {
	ans, err := s_help.Enriche(ctx, data.Name)
	if err != nil {
		return 0, err
	}

	fmt.Println(ans)

	return 0, nil
}
