package v1

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Response{message})
}
