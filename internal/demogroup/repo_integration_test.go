package demogroup_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/soadmized/banners_rotator/internal/config"
	"github.com/soadmized/banners_rotator/internal/demogroup"
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

	coll := db.Collection("demogroup-test")
	s.coll = coll

	// fixture
	_, err = s.coll.InsertOne(ctx, bson.M{"_id": "group13", "description": "desc group 13"})
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TearDownSuite() {
	ctx := context.Background()

	err := s.db.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestGet() {
	ctx := context.Background()
	repo := demogroup.Repo{Collection: s.coll}

	got, err := repo.Get(ctx, "group13")
	s.Require().NoError(err)
	s.Require().Equal(demogroup.Group{ID: "group13", Description: "desc group 13"}, *got)
}

func (s *RepoTestSuite) TestGetNotFound() {
	ctx := context.Background()
	repo := demogroup.Repo{Collection: s.coll}

	_, err := repo.Get(ctx, "group1111")
	s.Require().Error(err)
}

func (s *RepoTestSuite) TestCreate() {
	ctx := context.Background()
	repo := demogroup.Repo{Collection: s.coll}

	err := repo.Create(ctx, "group1", "desc group 1")
	s.Require().NoError(err)

	got, err := repo.Get(ctx, "group1")
	s.Require().NoError(err)
	s.Require().Equal(demogroup.Group{ID: "group1", Description: "desc group 1"}, *got)
}

func (s *RepoTestSuite) TestCreateExist() {
	ctx := context.Background()
	repo := demogroup.Repo{Collection: s.coll}

	err := repo.Create(ctx, "group13", "desc group 13")
	s.Require().Error(err)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
