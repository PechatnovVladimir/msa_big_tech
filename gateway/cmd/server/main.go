package main

import (
	"github.com/PechatnovVladimir/msa_big_tech/gateway/internal/gateway"
	"log"
)

func main() {
	log.Println("starting gateway...")
	go gateway.Rest()

	for {
		_ = 1
	}
}
