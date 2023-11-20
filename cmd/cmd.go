package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/build"
	"github.com/soadmized/banners_rotator/internal/config"
)

func Run(ctx context.Context, conf config.Config) error {
	builder, err := build.New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "build app")
	}

	api := builder.API()
	addr := fmt.Sprintf(":%d", conf.AppPort)

	log.Fatal(api.Router.Run(addr))

	return nil
}
