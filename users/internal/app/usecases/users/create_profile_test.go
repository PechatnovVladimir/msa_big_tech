package users

import (
	"context"
	"fmt"
	uRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/inmemory/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"testing"
)

func TestUserService_CreateProfile(t *testing.T) {
	ctx := context.Background()
	userRepo := uRepo.New(5)
	userService := New(userRepo)

	dtoProfile := dto.CreateProfileDTO{
		Nickname: "pvv45",
		Bio:      "pvvBiography",
		Avatar:   "pvv.jpeg",
		Password: "pvvPassword",
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

	u, err = userService.GetProfileByID(ctx, u.ID.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u)

	d := dto.UpdateProfileDTO{
		ID:     u.ID.String(),
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
