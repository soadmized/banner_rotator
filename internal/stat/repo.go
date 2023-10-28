package stat

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Collection *mongo.Collection
}

type slotStatDoc struct {
	ID         string        `bson:"_id"`
	BannerStat bannerStatDoc `bson:"banner_stat"`
}

type bannerStatDoc map[string]groupStatDoc

type groupStatDoc map[string]statDoc

type statDoc struct {
	Clicks int `bson:"clicks"`
	Shows  int `bson:"shows"`
}

func docToStat(d statDoc) Stat {
	return Stat{
		Clicks: d.Clicks,
		Shows:  d.Shows,
	}
}

func (r *Repo) GetStat(ctx context.Context, slotID, bannerID, groupID string) (*Stat, error) {
	var slotDoc slotStatDoc

	err := r.Collection.FindOne(ctx, bson.M{"_id": slotID}).Decode(&slotDoc)
	if err != nil {
		return nil, err
	}

	stat := slotDoc.BannerStat[bannerID][groupID]
	statModel := docToStat(stat)

	return &statModel, nil
}

func (r *Repo) AddClick(ctx context.Context, slotID, bannerID, groupID string) error {
	filter := bson.M{"_id": slotID}
	path := fmt.Sprintf("banner_stat.%s.%s.clicks", bannerID, groupID)
	set := bson.M{"$inc": bson.M{path: 1}}

	_, err := r.Collection.UpdateOne(ctx, filter, set, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "add click")
	}

	return nil
}

func (r *Repo) AddShow(ctx context.Context, slotID, bannerID, groupID string) error {
	filter := bson.M{"_id": slotID}
	path := fmt.Sprintf("banner_stat.%s.%s.shows", bannerID, groupID)
	set := bson.M{"$inc": bson.M{path: 1}}

	_, err := r.Collection.UpdateOne(ctx, filter, set, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "add show")
	}

	return nil
}
