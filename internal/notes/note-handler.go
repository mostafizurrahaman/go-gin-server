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

func (h *Handler) GetNoteList(c *gin.Context) {

	// Extract query:

	var query FilterParams

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	notes, err := h.repo.GetNoteList(c.Request.Context(), query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if notes == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "All notes retrived successfully.",
			"data":    []Note{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "All notes retrived successfully.",
		"data":    notes,
	})

}
