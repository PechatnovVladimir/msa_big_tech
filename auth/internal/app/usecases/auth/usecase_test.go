package auth

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/repositories/auth/inmemory"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"testing"
)

func TestService_Register(t *testing.T) {
	ctx := context.Background()
	authRepo := inmemory.NewInMemory()
	authService := New(Deps{
		AuthRepo:    authRepo,
		UserService: nil,
	})

	inRegister := dto.RegisterInDTO{
		Email:    "pvv@mail.ru",
		Password: "1234567890",
	}

	out, err := authService.Register(ctx, inRegister)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Register....", out.UserID)

	inLogin := dto.LoginInDTO{
		Email:    inRegister.Email,
		Password: inRegister.Password,
	}
	outLogin, err := authService.Login(ctx, inLogin)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("Login...", outLogin.UserID, outLogin.AccessToken, outLogin.RefreshToken)

	inRefresh := dto.RefreshInDTO{
		RefreshToken: outLogin.RefreshToken,
	}

	outRefresh, err := authService.Refresh(ctx, inRefresh)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Refresh...", outRefresh.UserID, outRefresh.AccessToken, outRefresh.RefreshToken)
}
