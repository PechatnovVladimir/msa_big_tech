package boot

import "google.golang.org/grpc"

func (app *App) RegisterGRPC(registerFunc func(*grpc.Server)) {
	app.grpcRegister = registerFunc
}
