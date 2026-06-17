package notes

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func getSortingValue(sortOrder TSortOrder) int {
	fmt.Println("sortOrder", sortOrder)
	orderStr := strings.ToUpper(string(sortOrder))
	if orderStr == "ASC" {
		return 1
	}

	return -1

}

func (r *Repo) GetNoteList(ctx context.Context, query FilterParams) ([]Note, error) {

	optCtx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	// ? Create a filter :
	filter := bson.M{}

	if query.FromDate != nil || query.ToDate != nil {
		dateFilter := bson.M{}

		if query.FromDate != nil {
			dateFilter["$gte"] = *query.FromDate
		}
		if query.ToDate != nil {
			endDate := query.ToDate.Add(24 * time.Hour).Add(-1 * time.Second)
			dateFilter["$lte"] = endDate
		}

		if len(dateFilter) > 0 {
			filter["creaedAt"] = dateFilter
		}
	}

	sorting := bson.D{{
		Key:   "createdAt",
		Value: -1,
	}}

	if query.SortBy != "" {
		sorting = bson.D{{
			Key:   query.SortBy,
			Value: getSortingValue(query.SortOrder),
		}}
	}

	fmt.Println(sorting)

	skip := (query.Page - 1) * query.Limit

	filterOptions := options.Find().SetSort(sorting).SetSkip(int64(skip)).SetLimit(int64(query.Limit))

	cursor, err := r.coll.Find(optCtx, filter, filterOptions)

	if err != nil {
		return nil, fmt.Errorf("Failed to read note list.")
	}

	defer cursor.Close(optCtx)

	var result []Note

	if err := cursor.All(optCtx, &result); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return result, nil

}
