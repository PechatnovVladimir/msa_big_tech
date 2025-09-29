package users

import (
	"context"
	"fmt"
	uRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/inmemory/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"github.com/google/uuid"
	"testing"
)

func TestUserService_CreateProfile(t *testing.T) {
	ctx := context.Background()
	userRepo := uRepo.New(5)
	userService := New(userRepo)

	dtoProfile := dto.CreateProfileDTO{
		ID:       uuid.New().String(),
		Nickname: "pvv45",
		Bio:      "pvvBiography",
		Avatar:   "pvv.jpeg",
	}

	u, err := userService.CreateProfile(ctx, dtoProfile)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u)

	_, err = userService.CreateProfile(ctx, dtoProfile)
	if err != nil {
		fmt.Println(err.Error())
	}

	u, err = userService.GetProfileByID(ctx, u.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u)

	d := dto.UpdateProfileDTO{
		ID:     u.ID,
		Bio:    "pvvBiographyNEW",
		Avatar: "pvvNEW.jpeg",
	}

	u, err = userService.UpdateProfile(ctx, d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u)

}
