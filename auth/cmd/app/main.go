package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app"
	"log"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatalln("auth app.Run error:", err)
	}
}
