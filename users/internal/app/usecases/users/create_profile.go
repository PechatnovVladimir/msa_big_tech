package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"regexp"
)

var (
	ErrCreateProfileFailed = errors.New("create profile failed")
)

func (s *Service) CreateProfile(ctx context.Context, d dto.CreateProfileDTO) (*users.UserProfile, error) {
	//валидация dto(?)

	//проверяем соответствие nickname маске ^[a-z0-9_]{3,20}$
	pattern := "^[a-z0-9_]{3,20}$"
	re := regexp.MustCompile(pattern)
	if !re.MatchString(d.Nickname) {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserInvalidArgument.Error())
	}

	//проверяем уникальность nickname
	_, err := s.GetProfileByNickname(ctx, d.Nickname)
	if err == nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	}

	//создаем профайл
	userProfile := users.NewUserProfile()
	//кеширование пароля
	psw := cachePassword(d.Password)

	userProfile.Nickname = d.Nickname
	userProfile.Bio = d.Bio
	userProfile.Avatar = d.Avatar
	userProfile.Password = psw

	err = s.repository.CreateProfile(ctx, userProfile)

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, err.Error())
	}
	return userProfile, nil
}
