package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/adapters/userprovider"
	serv "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc/v1"
	repo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	uc "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	pb "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc"
)

func Start(ctx context.Context) (err error) {

	app, err := boot.NewApp(ctx,
		boot.WithConfig(ctx, "config.yaml"),
	)

	if err != nil {
		return err
	}

	conn, txManager, err := app.Postgres(ctx)

	if err != nil {
		return err
	}
	defer conn.Close()

	userProvider := userprovider.New()

	repository := repo.New(txManager)

	userUseCase := uc.New(uc.Deps{
		UserRepo:     repository,
		UserProvider: userProvider,
	})

	service := serv.New(userUseCase)

	app.RegisterGRPC(func(srv *grpc.Server) {
		pb.RegisterUserServiceServer(srv, service)
	})

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
