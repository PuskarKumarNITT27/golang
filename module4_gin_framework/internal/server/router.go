package server

import (
	"gin-crud/internal/notes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func NewRouter(database *mongo.Database) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"success":true,
			"status":"Healthy",
		})
	})

	notes.RegisterRoutes(r,database)
	return r
}