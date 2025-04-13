package rest

import (
	"calculator/internal/pkg/model"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getCalculateResponse struct {
	Items []*model.Result `json:"items"`
}
