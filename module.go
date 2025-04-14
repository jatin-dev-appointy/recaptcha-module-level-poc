package main

import (
	"dummy_1"
	"dummy_2"
	"go.saastack.io/deployment/config"
	"go.uber.org/fx"
	"recaptcha_module"
)

var Module = fx.Options(
	fx.Provide(
		NewServeMux,   // Provides the HTTP mux for routing.
		NewHTTPClient, // Provides an HTTP client.
		NewLogger,     // Provides a logger for application logs
		recaptcha_module.NewReCaptchaServer,
	),
	dummy_1.Module,
	dummy_2.Module,
	config.Module,
	fx.Invoke(RunHTTPServer),
)
