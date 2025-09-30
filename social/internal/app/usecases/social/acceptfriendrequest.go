package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, in dto.AcceptFriendRequestIN) (dto.AcceptFriendRequestOUT, error) {
	log.Println("AcceptFriendRequest")

	//тестовый поход в репозиторий
	err := s.SocialRepo.Test()
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	//тестовый поход в сервис UserService
	err = s.UserService.Test()
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	return dto.AcceptFriendRequestOUT{}, nil
}
