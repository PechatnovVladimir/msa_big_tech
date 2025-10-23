package main

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/lib/secrets"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	//секреты в env
	//secret := secrets.NewSecrets(secrets.NewEnvProvider())

	secret := secrets.NewSecrets(secrets.NewFileProvider("secrets.yaml"))

	cfg, err := config.LoadConfig(ctx, secret)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("config: %+v", cfg)

	err = app.Run(ctx, cfg)
	if err != nil {
		log.Fatalln("app.Run error:", err)
	}
}
