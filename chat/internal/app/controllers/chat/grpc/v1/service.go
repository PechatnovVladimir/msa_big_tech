package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

// зависимости
type Deps struct {
	ChatUseCase chat.UseCase
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
