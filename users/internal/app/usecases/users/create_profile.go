package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"time"
)

var (
	ErrCreateProfileFailed = errors.New("create profile failed")
)

func (s *Service) CreateProfile(ctx context.Context, d dto.CreateProfileDTO) (*users.UserProfile, error) {
	//валидация dto (?)
	//валидацию сделали на уровне controllers, нужна ли еще дополнительная?

	//проверяем уникальность ID
	//_, err := s.GetProfileByID(ctx, d.ID)
	//if err == nil {
	//	return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	//}

	//проверяем уникальность nickname
	//_, err = s.GetProfileByNickname(ctx, d.Nickname)
	//if err == nil {
	//	return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	//}

	//создаем профайл
	userProfile := users.NewUserProfile()

	userProfile.ID = d.ID
	userProfile.Email = "pvv@mail.ru"
	userProfile.Nickname = d.Nickname
	userProfile.Bio = d.Bio
	userProfile.Avatar = d.Avatar
	userProfile.CreateAt = time.Now()

	//сохраняем в репозиторий
	err := s.repository.CreateProfile(ctx, userProfile)

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, err.Error())
	}
	return userProfile, nil
}
