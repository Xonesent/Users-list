package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error_Response struct {
	Message string `json:"message"`
}

// type statusResponse struct {
// 	Status string `json:"status"`
// }

func New_Error_Response(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Error_Response{message})
}
