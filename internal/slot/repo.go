package slot

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	repo *mongo.Collection
}

type slotDoc struct {
	ID          string `bson:"_id"`
	Description string `bson:"description"`
}

func docToSlot(d slotDoc) *Slot {
	return &Slot{
		ID:          ID(d.ID),
		Description: d.Description,
	}
}

func (r *Repo) Get(ctx context.Context, id ID) (*Slot, error) {
	var slot slotDoc

	err := r.repo.FindOne(ctx, bson.M{"_id": id}).Decode(&slot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("slot not found")
		}

		return nil, errors.Wrap(err, "get slot")
	}

	return docToSlot(slot), nil
}

func (r *Repo) Create(ctx context.Context, id ID, desc string) error {
	_, err := r.repo.InsertOne(ctx, bson.M{"_id": id, "description": desc})
	if err != nil {
		return errors.Wrap(err, "create slot")
	}

	return nil
}
