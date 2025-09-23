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
	usersRepo := uRepo.NewRepository(5)
	usersService := NewUserService(usersRepo)

	dtoProfile := dto.CreateProfileDTO{
		Nickname: "pvv45",
		Bio:      "pvvBiography",
		Avatar:   "pvv.jpeg",
		Password: "pvvPassword",
	}

	u, err := usersService.CreateProfile(ctx, dtoProfile)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u)

	_, err = usersService.CreateProfile(ctx, dtoProfile)
	if err != nil {
		fmt.Println(err.Error())
	}

	u, err = usersService.GetProfileByID(ctx, u.ID.String())
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

	u, err = usersService.UpdateProfile(ctx, d)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u)

}
