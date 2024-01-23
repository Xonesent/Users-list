package handler_api

import (
	"users-list/internal/service"

	"github.com/gin-gonic/gin"
)

type People_handler struct {
	people service.People
}

func New_People_handler(people service.People) *People_handler {
	return &People_handler{people: people}
}

func (h *People_handler) Get_Person(c *gin.Context) {

}

func (h *People_handler) Delete_Person(c *gin.Context) {

}

func (h *People_handler) Patch_Person(c *gin.Context) {

}

func (h *People_handler) Post_Person(c *gin.Context) {
	
}
