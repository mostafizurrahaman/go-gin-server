package notes

import "time"

type Note struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Pinned    bool      `json:"pinned" bson:"pinned"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type CreateNewNote struct {
	Title   string `json:"title" binding:"required,min=3,max=15"`
	Content string `json:"content" binding:"required,min=3"`
	Pinned  bool   `json:"pinned"`
}
