module recaptcha_module

go 1.23.4

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	go.saastack.io/deployment/config v0.0.0-20241203100103-2d2226f09c5f
	go.uber.org/fx v1.10.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.31.0
)

require (
	cloud.google.com/go v0.65.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20210106214847-113979e3529a // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.3 // indirect
)

replace (
	github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.14.6
	go.uber.org/dig => github.com/paullen/dig v1.7.1-0.20190624104937-6e47ebbbdcf6
	go.uber.org/fx => github.com/appointy/fx v1.9.1-0.20190624110333-490d04d33ef6
)
