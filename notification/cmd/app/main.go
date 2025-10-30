package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app"
	"log"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatalln("notification app.Run error:", err)
	}
}
