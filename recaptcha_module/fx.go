package recaptcha_module

import (
	"go.saastack.io/deployment/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(
	NewReCaptchaServer,
)

func NewReCaptchaServer(l config.Loader, logger *zap.Logger) (ReCaptchaServer, error) {
	def := ReCaptchaServer{
		ReCaptchaSkip: false,
		// test key, that will allow every token to pass
		ReCaptchaSecretKey: "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe",
	}

	if err := l.Load("recaptcha", &def); err != nil {
		return ReCaptchaServer{}, err
	}

	def.logger = logger
	return def, nil
}
