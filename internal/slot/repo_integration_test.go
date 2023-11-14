package slot_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/soadmized/banners_rotator/internal/config"
	"github.com/soadmized/banners_rotator/internal/slot"
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

	coll := db.Collection("slot-test")
	s.coll = coll

	// fixture
	_, err = s.coll.InsertOne(ctx, bson.M{"_id": "slot66", "description": "desc slot 66"})
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TearDownSuite() {
	ctx := context.Background()

	err := s.db.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestGet() {
	ctx := context.Background()
	repo := slot.Repo{Collection: s.coll}

	got, err := repo.Get(ctx, "slot66")
	s.Require().NoError(err)
	s.Require().Equal(slot.Slot{ID: "slot66", Description: "desc slot 66"}, *got)
}

func (s *RepoTestSuite) TestGetNotFound() {
	ctx := context.Background()
	repo := slot.Repo{Collection: s.coll}

	_, err := repo.Get(ctx, "slot777")
	s.Require().Error(err)
}

func (s *RepoTestSuite) TestCreate() {
	ctx := context.Background()
	repo := slot.Repo{Collection: s.coll}

	err := repo.Create(ctx, "slot1", "desc slot 1")
	s.Require().NoError(err)

	got, err := repo.Get(ctx, "slot1")
	s.Require().NoError(err)
	s.Require().Equal(slot.Slot{ID: "slot1", Description: "desc slot 1"}, *got)
}

func (s *RepoTestSuite) TestCreateExist() {
	ctx := context.Background()
	repo := slot.Repo{Collection: s.coll}

	err := repo.Create(ctx, "slot66", "desc slot 66")
	s.Require().Error(err)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
