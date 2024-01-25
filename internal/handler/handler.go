package handler

import (
	h_api "users-list/internal/handler/api"
	"users-list/internal/service"

	"github.com/gin-gonic/gin"
)

func New_Handler(services *service.Service) *Handler {
	return &Handler{
		People: h_api.New_People_handler(services.People),
	}
}

func (h *Handler) Init_Routes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/person", h.Get_Person)
		api.DELETE("/person/:id", h.Delete_Person)
		api.PATCH("/person", h.Patch_Person)
		api.POST("/person", h.Post_Person)
	}

	return router
}

type People interface {
	Get_Person(c *gin.Context)
	Delete_Person(c *gin.Context)
	Patch_Person(c *gin.Context)
	Post_Person(c *gin.Context)
}

type Handler struct{
	People
}
