package main

import (
	"context"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"log"

	"github.com/PechatnovVladimir/msa_big_tech/lib/config"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	cfgLoader := config.NewConfigLoader()

	//cfg, err := cfgLoader.LoadConfig("/Users/pvv/Educ/Balun/msa_big_tech/chat/cmd/client/config.yaml")
	cfg, err := cfgLoader.LoadConfig("./chat/cmd/client/config.yaml")

	if err != nil {
		log.Println("Error loading config:", err)
		log.Fatal(err)
	}

	chat, err := NewClientNew("localhost:50052", cfg)
	if err != nil {
		log.Printf("Error creating client: %v", err)
		log.Fatal(err)

	}

	//defer chat.Close()

	userID1 := uuid.New().String()

	in := &chatPB.CreateDirectChatRequest{
		ParticipantId: userID1,
	}

	for range 1 {

		resp, err := chat.client.CreateDirectChat(ctx, in)
		if err != nil {
			logger.Fatal(ctx, err)
		}

		chatID1 := resp.ChatId

		for range 100 {
			message := &chatPB.SendMessageRequest{
				ChatId: chatID1,
				Text:   gofakeit.Comment(),
			}

			rsp, err := chat.client.SendMessage(ctx, message)

			if err != nil {
				logger.Fatal(ctx, err)
			}
			log.Println(chatID1, rsp)
		}
	}
}
