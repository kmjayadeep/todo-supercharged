package todo

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
	CompletedAt time.Time          `json:"completedAt"`
	Completed   bool               `json:"completed"`
}

type TodoController struct {
	Db *mongo.Database
}

func (tc *TodoController) GetTodos(c *gin.Context) {
	var todos []Todo
	ctx := c.Request.Context()

	collection := tc.Db.Collection("todo")

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = cur.All(ctx, &todos)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, todos)

}

func (tc *TodoController) AddTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.CompletedAt = time.Now()
	todo.Completed = false

	collection := tc.Db.Collection("todo")

	doc, err := collection.InsertOne(c.Request.Context(), todo)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, doc)
}
