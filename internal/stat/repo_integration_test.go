package stat_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/soadmized/banners_rotator/internal/banner"
	"github.com/soadmized/banners_rotator/internal/config"
	"github.com/soadmized/banners_rotator/internal/stat"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoTestSuite struct {
	suite.Suite
	conn *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

func (s *RepoTestSuite) SetupSuite() {
	ctx := context.Background()
	conf := config.Load()

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI(conf)))
	s.Require().NoError(err)
	s.conn = conn

	db := conn.Database(fmt.Sprintf("test-%s", uuid.New().String()))
	s.db = db
}

func (s *RepoTestSuite) TearDownSuite() {
	ctx := context.Background()

	err := s.db.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) SetupTest() {
	ctx := context.Background()

	coll := s.db.Collection("stat-test")
	s.coll = coll

	// fixture
	_, err := s.coll.InsertOne(ctx, bson.M{"_id": "slot777", "banner_stat": bson.M{
		"banner1": bson.M{"group1": bson.M{"clicks": 1, "shows": 1}, "group2": bson.M{"clicks": 2, "shows": 2}},
	}})
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TearDownTest() {
	ctx := context.Background()

	err := s.coll.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestGetStat() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	got, err := repo.GetStat(ctx, "slot777", "banner1", "group2")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 2, Shows: 2}, *got)
}

func (s *RepoTestSuite) TestGetStatSlotNotExist() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	got, err := repo.GetStat(ctx, "slot098", "banner1", "group2")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 0, Shows: 0}, *got)
}

func (s *RepoTestSuite) TestGetStatGroupNotExist() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	got, err := repo.GetStat(ctx, "slot777", "banner1", "group234")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 0, Shows: 0}, *got)
}

func (s *RepoTestSuite) TestGetStatBannerNotExist() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	got, err := repo.GetStat(ctx, "slot777", "banner42", "group1")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 0, Shows: 0}, *got)
}

func (s *RepoTestSuite) TestGetBannerIDs() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	got, err := repo.GetBannerIDs(ctx, "slot777")
	s.Require().NoError(err)
	s.Require().Equal([]banner.ID{"banner1"}, got)
}

func (s *RepoTestSuite) TestAddClick() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	err := repo.AddClick(ctx, "slot777", "banner2", "group2")
	s.Require().NoError(err)

	got, err := repo.GetStat(ctx, "slot777", "banner2", "group2")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 1, Shows: 0}, *got)
}

func (s *RepoTestSuite) TestAddShow() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	err := repo.AddShow(ctx, "slot777", "banner3", "group2")
	s.Require().NoError(err)

	got, err := repo.GetStat(ctx, "slot777", "banner3", "group2")
	s.Require().NoError(err)
	s.Require().Equal(stat.Stat{Clicks: 0, Shows: 1}, *got)
}

func (s *RepoTestSuite) TestAddBanner() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	err := repo.AddBanner(ctx, "slot666", "banner777")
	s.Require().NoError(err)

	ids, err := repo.GetBannerIDs(ctx, "slot666")
	s.Require().NoError(err)
	s.Require().Equal([]banner.ID{"banner777"}, ids)
}

func (s *RepoTestSuite) TestAddExistingBanner() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	err := repo.AddBanner(ctx, "slot777", "banner1")
	s.Require().Error(err)
}

func (s *RepoTestSuite) TestRemoveBanner() {
	ctx := context.Background()
	repo := stat.Repo{Collection: s.coll}

	// remove banner from fixture slot
	err := repo.RemoveBanner(ctx, "slot777", "banner1")
	s.Require().NoError(err)

	ids, err := repo.GetBannerIDs(ctx, "slot777")
	s.Require().NoError(err)
	s.Require().Equal([]banner.ID{}, ids)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
