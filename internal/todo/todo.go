package todo

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetTodos(c *gin.Context) {
}
