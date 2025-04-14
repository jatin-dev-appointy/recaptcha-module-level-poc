package dummy_2

import (
	"context"
	"dummy_2/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDummy2Server,

		fx.Annotated{
			Name: "public",
			Target: func(srv pb.Dummy2Server) pb.Dummy2Server {
				return srv
			},
		},

		fx.Annotated{
			Name: "public",
			Target: func(in struct {
				fx.In
				S pb.Dummy2Server `name:"public"`
			}) pb.Dummy2Client {
				return pb.NewLocalDummy2Client(in.S)
			},
		},

		fx.Annotated{
			Group:  "http-service",
			Target: RegisterDummy2HttpService,
		},
	),

	pb.UnregisteredMethods_Dummy2,
	fx.Decorate(
		pb.NewDummy2RecaptchaServer,
	),
)

func RegisterDummy2HttpService(in struct {
	fx.In
	Client pb.Dummy2Client `name:"public"`
}) func(*runtime.ServeMux, context.Context) error {
	return func(mux *runtime.ServeMux, ctx context.Context) error {
		return pb.RegisterDummy2HandlerClient(ctx, mux, in.Client)
	}
}
