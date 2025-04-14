package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"net/http"
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func NewServeMux() *runtime.ServeMux {
	return runtime.NewServeMux()
}

type In struct {
	fx.In
	ServeMux      *runtime.ServeMux
	HttpRegisters []func(mux *runtime.ServeMux, ctx context.Context) error `group:"http-service"`
}

func RunHTTPServer(lc fx.Lifecycle, in In) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			// Add CORS middleware
			corsMiddleware := cors.New(cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}).Handler(in.ServeMux)

			// Register HTTP routes
			for _, register := range in.HttpRegisters {
				if err := register(in.ServeMux, ctx); err != nil {
					return err
				}
			}

			// add custom middlewares
			handler := SetGoogleCaptchaResponseInContext(corsMiddleware)

			// Start HTTP server with CORS middleware
			go func() {
				if err := http.ListenAndServe(":8080", handler); err != nil {
					fmt.Printf("HTTP server failed: %v\n", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server...")
			return nil
		},
	})
}
