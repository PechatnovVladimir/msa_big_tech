package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app"
)

func main() {
	ctx := context.Background()

	err := app.Start(ctx)
	if err != nil {
		logger.Fatal(ctx, "app.Run error:", err)
	}
}
