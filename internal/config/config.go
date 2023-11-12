package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort int32 `envconfig:"APP_PORT" default:"8080"`

	MongoHost          string `envconfig:"MONGO_HOST" default:"localhost"`
	MongoDBName        string `envconfig:"MONGO_DB"`
	MongoBannerColl    string `envconfig:"MONGO_BANNER_COLL"`
	MongoSlotColl      string `envconfig:"MONGO_SLOT_COLL"`
	MongoDemoGroupColl string `envconfig:"MONGO_DEMOGROUP_COLL"`
	MongoStatColl      string `envconfig:"MONGO_STAT_COLL"`
}

func Load() Config {
	conf := Config{}

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}

	return conf
}

func MongoURI(conf Config) string {
	uri := fmt.Sprintf("mongodb://%s", conf.MongoHost)

	return uri
}
