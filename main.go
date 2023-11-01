package main

import (
	"banners_rotator/cmd"
	"banners_rotator/internal/config"
	"context"
	"log"
)

func main() {
	conf := config.Load()
	ctx := context.Background()

	err := cmd.Run(ctx, conf)
	if err != nil {
		log.Print(err)
	}
}
