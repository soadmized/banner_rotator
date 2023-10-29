package cmd

import (
	"banners_rotator/internal/build"
	"banners_rotator/internal/config"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func Run(ctx context.Context, conf config.Config) error {
	builder, err := build.New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "build app")
	}

	api := builder.Api()
	addr := fmt.Sprintf("localhost:%d", conf.AppPort)

	log.Fatal(api.Router.Run(addr))

	return nil
}
