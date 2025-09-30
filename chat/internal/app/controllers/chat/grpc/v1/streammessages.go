package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc"
)

func (s *Service) StreamMessages(*chat.StreamMessagesRequest, grpc.ServerStreamingServer[chat.StreamMessagesResponse]) error {
	return nil
}
