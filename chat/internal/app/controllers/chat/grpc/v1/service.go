package v1

import (
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

// зависимости
type Deps struct {
}

type Service struct {
	chatPB.UnimplementedChatServiceServer
	Deps
}

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
