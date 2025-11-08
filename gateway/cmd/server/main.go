package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/gateway/internal/gateway"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func main() {
	logger.Info(context.TODO(), "starting gateway...")
	go gateway.Rest()

	for {
		_ = 1
	}
}
