package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	todo.ID = primitive.NewObjectID()
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

func (tc *TodoController) DeleteTodo(c *gin.Context) {

	id := c.Param("id")
	idObj, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	collection := tc.Db.Collection("todo")

	filter := bson.D{
		{Key: "_id", Value: idObj},
	}

	res, err := collection.DeleteOne(c.Request.Context(), filter)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if res.DeletedCount == 1 {
		c.String(http.StatusOK, "Deleted successfully")
	} else {
		c.String(http.StatusNotFound, "Unable to find the todo")
	}
}

func (tc *TodoController) MarkDone(c *gin.Context) {

	id := c.Param("id")
	idObj, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	log.Info("deleting todo with id ", id)

	collection := tc.Db.Collection("todo")

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "completed", Value: true},
			{Key: "completedAt", Value: time.Now()},
		}},
	}

	res, err := collection.UpdateByID(c.Request.Context(), idObj, update)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if res.ModifiedCount == 1 {
		c.String(http.StatusOK, "Updated successfully")
	} else {
		c.String(http.StatusNotFound, "Couldn't update todo")
	}
}
