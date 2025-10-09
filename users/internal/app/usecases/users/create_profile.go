package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

var (
	ErrCreateProfileFailed = errors.New("create profile failed")
)

func (s *Service) CreateProfile(ctx context.Context, profile dto.CreateProfile) (*users.UserProfile, error) {
	//проверяем уникальность ID
	_, err := s.GetProfileByID(ctx, dto.GetProfileById{ID: profile.ID})
	if err == nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	}

	//проверяем уникальность nickname
	_, err = s.GetProfileByNickname(ctx, dto.GetProfileByNickname{Nickname: profile.Nickname})
	if err == nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateProfileFailed, users.ErrUserAlreadyExists.Error())
	}

	data := modelUserProfileFromCreateProfileDto(profile)

	//сохраняем в репозиторий
	outProfile, err := s.repository.CreateProfile(ctx, data)

	if err != nil {
		return &users.UserProfile{}, fmt.Errorf("%w: %s", ErrCreateProfileFailed, err.Error())
	}
	return outProfile, nil
}
