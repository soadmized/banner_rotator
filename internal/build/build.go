package build

import (
	"banners_rotator/internal/banner"
	"banners_rotator/internal/config"
	"banners_rotator/internal/demogroup"
	"banners_rotator/internal/slot"
	"banners_rotator/internal/stat"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type BannerService interface {
	Get(ctx context.Context, id banner.ID) (*banner.Banner, error)
	Create(ctx context.Context, id banner.ID, desc string) error
}

type DemoGroupService interface {
	Get(ctx context.Context, id demogroup.ID) (*demogroup.Group, error)
	Create(ctx context.Context, id demogroup.ID, desc string) error
}

type SlotService interface {
	Get(ctx context.Context, id slot.ID) (*slot.Slot, error)
	Create(ctx context.Context, id slot.ID, desc string) error
}

type StatService interface {
	GetStat(ctx context.Context, slotID, bannerID, groupID string) (*stat.Stat, error)
	AddShow(ctx context.Context, slotID, bannerID, groupID string) error
	AddClick(ctx context.Context, slotID, bannerID, groupID string) error
}

type Builder struct {
	conf config.Config

	mongoClient *mongo.Client
	mongoDB     mongo.Database
	collections map[string]mongo.Collection
}

func New(ctx context.Context, conf config.Config) *Builder {
	return nil
}
