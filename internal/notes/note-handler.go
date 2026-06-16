package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context) {

	// ? Read request body :
	var req CreateNewNote

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind the json!",
			"data":    err.Error(),
		})

		return
	}

	now := time.Now().UTC()

	// ? Configure new struct for note record:
	var note = Note{
		ID:        bson.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// ?? call the repo here to save into db (as like service)
	newNote, err := h.repo.CreateNote(c.Request.Context(), note)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to save note into db!",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": false,
		"message": "Note saved successfully into db!",
		"data":    newNote,
	})

}
