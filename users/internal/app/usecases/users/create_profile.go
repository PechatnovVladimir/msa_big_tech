package users

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

var (
	ErrCreateProfileFailed = errors.New("create profile failed")
)

func (s *UserService) CreateProfile(ctx context.Context, dto dto.CreateProfileDTO) (*users.UserProfile, error) {
	//валидация dto(?)

	//проверяем уникальность nickname
	_, err := s.GetProfileByNickname(ctx, dto.Nickname)
	if err == nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	}

	//создаем профайл
	userProfile := users.NewUserProfile()
	//кеширование пароля
	psw := cachePassword(dto.Password)

	userProfile.Nickname = dto.Nickname
	userProfile.Bio = dto.Bio
	userProfile.Avatar = dto.Avatar
	userProfile.Password = psw

	err = s.userRepo.CreateProfile(ctx, userProfile)

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, err.Error())
	}
	return userProfile, nil
}

func cachePassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
