package demogroup

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	repo *mongo.Collection
}

type groupDoc struct {
	ID          string `bson:"_id"`
	Description string `bson:"description"`
}

func docToGroup(d groupDoc) *Group {
	return &Group{
		ID:          ID(d.ID),
		Description: d.Description,
	}
}

func (r *Repo) Get(ctx context.Context, id ID) (*Group, error) {
	var group groupDoc

	err := r.repo.FindOne(ctx, bson.M{"_id": id}).Decode(&group)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("demo group not found")
		}

		return nil, errors.Wrap(err, "get demo group")
	}

	return docToGroup(group), nil
}

func (r *Repo) Create(ctx context.Context, id ID, desc string) error {
	_, err := r.repo.InsertOne(ctx, bson.M{"_id": id, "description": desc})
	if err != nil {
		return errors.Wrap(err, "create demo group")
	}

	return nil
}
