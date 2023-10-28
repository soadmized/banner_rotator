package build

import (
	"banners_rotator/internal/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Mongo *Mongo
}

type Mongo struct {
	BannerColl    *mongo.Collection
	SlotColl      *mongo.Collection
	DemoGroupColl *mongo.Collection
	StatsColl     *mongo.Collection
}

func Build(ctx context.Context, conf config.Config) (*App, error) {
	mng, err := buildMongo(ctx, conf)
	if err != nil {
		return nil, err
	}

	return &App{Mongo: mng}, nil
}

func buildMongo(ctx context.Context, conf config.Config) (*Mongo, error) {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI(conf)))
	if err != nil {
		return nil, err
	}

	db := mongoClient.Database(conf.MongoDB)
	bannerColl := db.Collection(conf.MongoBannerColl)
	slotColl := db.Collection(conf.MongoSlotColl)
	groupColl := db.Collection(conf.MongoDemoGroupColl)
	statsColl := db.Collection(conf.MongoStatsColl)

	return &Mongo{
		BannerColl:    bannerColl,
		SlotColl:      slotColl,
		DemoGroupColl: groupColl,
		StatsColl:     statsColl,
	}, nil
}
