package main

import (
	"context"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"log"
)

func main() {
	ctx := context.Background()

	chat, err := NewClient("localhost:50052")
	if err != nil {
		log.Fatal(err)
	}
	defer chat.Close()

	userID1 := uuid.New().String()

	in := &chatPB.CreateDirectChatRequest{
		ParticipantId: userID1,
	}

	for range 1 {

		resp, err := chat.client.CreateDirectChat(ctx, in)
		if err != nil {
			log.Fatal(err)
		}

		chatID1 := resp.ChatId

		for range 100 {
			message := &chatPB.SendMessageRequest{
				ChatId: chatID1,
				Text:   gofakeit.Comment(),
			}

			rsp, err := chat.client.SendMessage(ctx, message)

			if err != nil {
				log.Fatal(err)
			}
			log.Println(chatID1, rsp)
		}
	}
}
