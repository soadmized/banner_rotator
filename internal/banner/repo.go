package banner

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Collection *mongo.Collection
}

type bannerDoc struct {
	ID          string `bson:"_id"`
	Description string `bson:"description"`
}

func docToBanner(d bannerDoc) *Banner {
	return &Banner{
		ID:          ID(d.ID),
		Description: d.Description,
	}
}

func (r *Repo) Get(ctx context.Context, id ID) (*Banner, error) {
	var banner bannerDoc

	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&banner)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("banner not found")
		}

		return nil, errors.Wrap(err, "get banner")
	}

	return docToBanner(banner), nil
}

func (r *Repo) Create(ctx context.Context, id ID, desc string) error {
	_, err := r.Collection.InsertOne(ctx, bson.M{"_id": id, "description": desc})
	if err != nil {
		return errors.Wrap(err, "create banner")
	}

	return nil
}
