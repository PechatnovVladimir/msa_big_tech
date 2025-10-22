package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) StreamMessages(ctx context.Context, in *dto.StreamMessagesIN) (<-chan *chat.Message, error) {
	const api = "ChatService.StreamMessages"

	const (
		chanelBufferSize    = 5
		messageHistoryLimit = 50
	)

	//проверка, а существует ли чат
	_, err := s.GetChat(ctx, &dto.GetChatIN{ChatID: in.ChatID})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	messageChanel := make(chan *chat.Message, chanelBufferSize)

	go func() {
		defer close(messageChanel)
		if in.SinceMessageTime.Unix() != 0 {
			_, paginationOpts := fromListMessagesIN(&dto.ListMessagesIN{
				ChatID: in.ChatID,
				Limit:  uint64(messageHistoryLimit),
				Cursor: dto.Cursor{
					Time: in.SinceMessageTime,
				},
			})
			res, err := s.ChatRepo.ListMessages(ctx, in.ChatID, paginationOpts)
			if err != nil {
				log.Printf("%s: GetMessages error: %s", api, err)
				return
			}
			for _, msg := range res {
				select {
				case messageChanel <- msg:
				case <-ctx.Done():
					return
				}
			}
		}
		streamChanel, errStream := s.ChatRepo.StreamMessage(ctx, in.ChatID)
		if errStream != nil {
			log.Printf("%s: StreamMessage error: %s", api, errStream)
			return
		}

		for {
			select {
			case msg, ok := <-streamChanel:
				if !ok {
					return
				}
				select {
				case messageChanel <- msg:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}

	}()

	return messageChanel, nil
}
