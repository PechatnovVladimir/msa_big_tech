package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"github.com/google/uuid"
	"log"
)

func (s *Service) Register(ctx context.Context, in dto.RegisterInDTO) (out dto.RegisterOutDTO, err error) {
	err = s.AuthRepo.CheckUserByEmail(ctx, in.Email)
	if err != nil {
		//пользователь уже существует
		log.Println("[ERROR] checking user by email: ")
		return dto.RegisterOutDTO{}, auth.ErrAuthAlreadyExists
	}

	userID := uuid.New().String()
	email := in.Email
	password := cachePassword(in.Password)

	err = s.AuthRepo.SaveUser(ctx, userID, email, password)
	if err != nil {
		return dto.RegisterOutDTO{}, auth.ErrAuthAlreadyExists
	}

	out.UserID = userID

	//надо создать профиль пользователя в сервисе пользователей

	return out, nil
}
