package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Server is running on port 5000",
			"data":    nil,
		})

	})

	return r

}
