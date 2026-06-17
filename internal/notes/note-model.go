package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Note struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content" bson:"content"`
	Pinned    bool          `json:"pinned" bson:"pinned"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

type CreateNewNote struct {
	Title   string `json:"title" binding:"required,min=3,max=15"`
	Content string `json:"content" binding:"required,min=3"`
	Pinned  bool   `json:"pinned"`
}

type TSortOrder string

const (
	ASC  TSortOrder = "asc"
	DESC TSortOrder = "desc"
)

type FilterParams struct {
	Page       int        `form:"page" binding:"required,min=1"`
	Limit      int        `form:"limit" binding:"required,min=1"`
	SearchTerm string     `form:"searchTerm" `
	FromDate   *time.Time `form:"fromDate"`
	ToDate     *time.Time `form:"toDate"`
	SortBy     string     `form:"sortBy"`
	SortOrder  TSortOrder `form:"sortOrder"`
}
