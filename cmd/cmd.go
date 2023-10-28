package cmd

import (
	"banners_rotator/internal/build"
	"banners_rotator/internal/config"
	"context"
)

func Run(ctx context.Context, conf config.Config) error {
	build.New(ctx, conf)

	return nil
}
