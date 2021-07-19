package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	// "go.mongodb.org/mongo-driver/bson"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var res string
	collection := client.Database("testdb").Collection("people")
	err = collection.FindOne(context.Background(), nil).Decode(&res)

	app := gin.Default()

	app.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Todo supercharged",
		})
	})

	app.Run()

}
