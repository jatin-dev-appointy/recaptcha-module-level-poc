package dummy_1

import (
	"context"
	"dummy_1/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDummy1Server,

		fx.Annotated{
			Name: "public",
			Target: func(srv pb.Dummy1Server) pb.Dummy1Server {
				return srv
			},
		},

		fx.Annotated{
			Name: "public",
			Target: func(in struct {
				fx.In
				S pb.Dummy1Server `name:"public"`
			}) pb.Dummy1Client {
				return pb.NewLocalDummy1Client(in.S)
			},
		},

		fx.Annotated{
			Group:  "http-service",
			Target: RegisterDummy1HttpService,
		},
	),

	pb.UnregisteredMethods_Dummy1,
	fx.Decorate(
		pb.NewDummy1RecaptchaServer,
	),
)

func RegisterDummy1HttpService(in struct {
	fx.In
	Client pb.Dummy1Client `name:"public"`
}) func(*runtime.ServeMux, context.Context) error {
	return func(mux *runtime.ServeMux, ctx context.Context) error {
		return pb.RegisterDummy1HandlerClient(ctx, mux, in.Client)
	}
}
