package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {

	return &Repo{
		coll: db.Collection("notes"),
	}

}

func (r *Repo) CreateNote(ctx context.Context, note Note) (*Note, error) {

	// ??  Operation context
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// ??  Insert data into db:
	_, err := r.coll.InsertOne(opCtx, note)

	if err != nil {
		return nil, fmt.Errorf("Failed to insert into db! %s", err.Error())
	}

	return &note, nil

}
