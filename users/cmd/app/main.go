package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	err := app.Start(ctx)
	if err != nil {
		log.Fatalln("app.Run error:", err)
	}
}
