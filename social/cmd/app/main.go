package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(context.Background(), cfg)
	if err != nil {
		log.Fatalln("social app.Run error:", err)
	}
}
