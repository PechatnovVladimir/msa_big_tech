package auth

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/google/uuid"
)

func (s *Service) Register(ctx context.Context, in dto.RegisterInDTO) (out dto.RegisterOutDTO, err error) {
	err = s.AuthRepo.CheckUserByEmail(ctx, in.Email)
	if err != nil {
		//пользователь уже существует
		logger.Error(ctx, "checking user by email: ")
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

	createProfileIN := dto.CreateUserProfileInDTO{
		UserID: out.UserID,
		//где вот тут мы возьмем значение nickname
		Nickname: "empty",
	}

	//надо создать профиль пользователя в сервисе пользователей
	_, err = s.UserService.CreateUserProfile(ctx, createProfileIN)
	if err != nil {
		logger.Error(ctx, "creating user: ")
		return dto.RegisterOutDTO{}, errors.New("error creating UserProfile in User Service")
	}

	return out, nil
}
