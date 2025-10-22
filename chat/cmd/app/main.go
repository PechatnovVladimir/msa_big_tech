package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("config: %+v", cfg)

	err = app.Run(context.Background(), cfg)
	if err != nil {
		log.Fatalln("chat app.Run error:", err)
	}
}
