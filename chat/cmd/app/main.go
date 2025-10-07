package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app"
	"log"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatalln("chat app.Run error:", err)
	}
}
