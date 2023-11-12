package build

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/soadmized/banners_rotator/internal/api"
	"github.com/soadmized/banners_rotator/internal/banner"
	"github.com/soadmized/banners_rotator/internal/config"
	"github.com/soadmized/banners_rotator/internal/demogroup"
	"github.com/soadmized/banners_rotator/internal/slot"
	"github.com/soadmized/banners_rotator/internal/stat"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Builder struct {
	conf config.Config

	mongoDB *mongo.Database
}

func New(ctx context.Context, conf config.Config) (*Builder, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI(conf)))
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.MongoDBName)

	builder := Builder{
		conf:    conf,
		mongoDB: db,
	}

	return &builder, nil
}

func (b *Builder) Api() *api.Api {
	router := gin.New()
	bannerSrv := b.bannerService()
	groupSrv := b.demoGroupService()
	slotSrv := b.slotService()
	statSrv := b.statService()

	api2 := api.Api{
		Router:    router,
		BannerSrv: &bannerSrv,
		SlotSrv:   &slotSrv,
		GroupSrv:  &groupSrv,
		StatSrv:   &statSrv,
	}

	api2.RegisterHandlers()

	return &api2
}

func (b *Builder) bannerService() banner.Service {
	repo := b.bannerRepo()
	srv := banner.Service{Repo: &repo}

	return srv
}

func (b *Builder) bannerRepo() banner.Repo {
	coll := b.mongoDB.Collection(b.conf.MongoBannerColl)

	repo := banner.Repo{Collection: coll}

	return repo
}

func (b *Builder) demoGroupService() demogroup.Service {
	repo := b.demoGroupRepo()
	srv := demogroup.Service{Repo: &repo}

	return srv
}

func (b *Builder) demoGroupRepo() demogroup.Repo {
	coll := b.mongoDB.Collection(b.conf.MongoDemoGroupColl)

	repo := demogroup.Repo{Collection: coll}

	return repo
}

func (b *Builder) slotService() slot.Service {
	repo := b.slotRepo()
	srv := slot.Service{Repo: &repo}

	return srv
}

func (b *Builder) slotRepo() slot.Repo {
	coll := b.mongoDB.Collection(b.conf.MongoSlotColl)

	repo := slot.Repo{Collection: coll}

	return repo
}

func (b *Builder) statService() stat.Service {
	repo := b.statRepo()
	srv := stat.Service{Repo: &repo}

	return srv
}

func (b *Builder) statRepo() stat.Repo {
	coll := b.mongoDB.Collection(b.conf.MongoStatColl)

	repo := stat.Repo{Collection: coll}

	return repo
}
