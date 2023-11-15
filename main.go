package main

import (
	"context"
	"log"

	"github.com/soadmized/banners_rotator/cmd"
	"github.com/soadmized/banners_rotator/internal/config"
)

func main() {
	conf := config.Load()
	ctx := context.Background()

	err := cmd.Run(ctx, conf)
	if err != nil {
		log.Print(err)
	}
}
