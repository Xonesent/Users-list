package h_api

import (
	"net/http"
	h_help "users-list/internal/handler/helpers"
	"users-list/internal/service"
	"users-list/server"

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
	var input server.Post_structure

	if err := c.BindJSON(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h_help.Validate_Post_Structure(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.people.Post_Person(c.Request.Context(), &input)
	if err != nil {
		server.New_Error_Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
