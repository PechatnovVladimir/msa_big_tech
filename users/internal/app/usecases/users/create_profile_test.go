package users

import (
	"context"
	"fmt"
	uRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"testing"
)

func TestUserService_CreateProfile(t *testing.T) {
	ctx := context.Background()
	usersRepo := uRepo.NewRepository()
	usersService := NewUserService(usersRepo)

	dtoProfile := dto.CreateProfileDTO{
		Nickname: "pvv",
		Bio:      "pvvBiography",
		Avatar:   "pvv.jpeg",
		Password: "pvvPassword",
	}

	u, err := usersService.CreateProfile(ctx, dtoProfile)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u)
}
