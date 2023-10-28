package cmd

import (
	"banners_rotator/internal/build"
	"banners_rotator/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func Run(ctx context.Context, conf config.Config) error {
	app, err := build.Build(ctx, conf)
	if err != nil {
		return err
	}

	app.Mongo.BannerColl.InsertOne(ctx, bson.M{"_id": "banner1", "desc": "first banner"})
	return nil
}
