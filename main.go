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

	app.StaticFile("/", "./client/build/index.html")
	app.Static("/static", "./client/build/static")

	app.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Todo supercharged",
		})
	})

	app.Use(CORSMiddleware())

	app.GET("/v1/todo", todoController.GetTodos)
	app.POST("/v1/todo", todoController.AddTodo)
	app.DELETE("/v1/todo/:id", todoController.DeleteTodo)
	app.PUT("/v1/todo/:id/done", todoController.MarkDone)

	app.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
