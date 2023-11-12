package stat

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/banner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	groupStat, ok := slotDoc.BannerStat[bannerID]
	if !ok {
		statModel := docToStat(statDoc{
			Clicks: 0,
			Shows:  0,
		})

		return &statModel, err
	}

	stat, ok := groupStat[groupID]
	if !ok {
		statModel := docToStat(statDoc{
			Clicks: 0,
			Shows:  0,
		})

		return &statModel, err
	}

	statModel := docToStat(stat)

	return &statModel, nil
}

func (r *Repo) GetBannerIDs(ctx context.Context, slotID string) ([]banner.ID, error) {
	var slotDoc slotStatDoc

	err := r.Collection.FindOne(ctx, bson.M{"_id": slotID}).Decode(&slotDoc)
	if err != nil {
		return nil, err
	}

	bannerIDs := make([]banner.ID, 0, len(slotDoc.BannerStat))

	for k := range slotDoc.BannerStat {
		id := banner.ID(k)
		bannerIDs = append(bannerIDs, id)
	}

	return bannerIDs, nil
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

func (r *Repo) AddBanner(ctx context.Context, slotID, bannerID string) error {
	filter := bson.M{"_id": slotID}
	path := fmt.Sprintf("banner_stat.%s", bannerID)
	set := bson.M{"$set": path}

	_, err := r.Collection.UpdateOne(ctx, filter, set, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "add banner")
	}

	return nil
}

func (r *Repo) RemoveBanner(ctx context.Context, slotID, bannerID string) error {
	filter := bson.M{"_id": slotID}
	path := fmt.Sprintf("banner_stat.%s", bannerID)
	set := bson.M{"$unset": path}

	_, err := r.Collection.UpdateOne(ctx, filter, set, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "remove banner")
	}

	return nil
}
