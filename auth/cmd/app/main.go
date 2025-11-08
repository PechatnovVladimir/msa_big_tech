package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func main() {
	ctx := context.Background()

	err := app.Start(ctx)
	if err != nil {
		logger.Fatalf(ctx, "app.Run error: %v", err)
	}
}
