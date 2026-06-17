package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterNoteRoutes(r *gin.Engine, db *mongo.Database) {

	// ? Create repo first :
	repo := NewRepo(db)

	// ? Create handler :
	handler := NewHandler(repo)

	noteGroup := r.Group("notes")

	{
		noteGroup.POST("/", handler.CreateNote)
		noteGroup.GET("/list", handler.GetNoteList)
	}

}
