package main

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"

	"github.com/kmjayadeep/todo-supercharged/internal/todo"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("connecting to mongodb with url", CONNECTION_STRING)

	connectOptions := options.Client().ApplyURI(CONNECTION_STRING)
	client, err := mongo.Connect(ctx, connectOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	log.Info("DB connected, starting server")

	db := client.Database("todo")
	todoController := todo.TodoController{
		Db: db,
	}

	app := gin.Default()

	app.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Todo supercharged",
		})
	})

	app.GET("/v1/todo", todoController.GetTodos)
	app.POST("/v1/todo", todoController.AddTodo)
	app.DELETE("/v1/todo/:id", todoController.DeleteTodo)

	app.Run()

}
