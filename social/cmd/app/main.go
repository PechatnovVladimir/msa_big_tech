package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app"
	"log"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatalln("social app.Run error:", err)
	}
}
