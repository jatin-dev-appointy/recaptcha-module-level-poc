// Code generated by recaptcha. DO NOT EDIT.

package pb

import (
	"context"
	"recaptcha_module"
)

// Service: Dummy2

type dummy2RecaptchaServer struct {
	Dummy2Server
	recaptchaServer recaptcha_module.ReCaptchaServer
}

func NewDummy2RecaptchaServer(srv Dummy2Server, rSrv recaptcha_module.ReCaptchaServer) Dummy2Server {
	return &dummy2RecaptchaServer{
		Dummy2Server:    srv,
		recaptchaServer: rSrv,
	}
}

// Method: GetDummy2Method1 (This method does not require reCAPTCHA verification)

// Method: GetDummy2Method2 (This method requires reCAPTCHA verification)
func (s *dummy2RecaptchaServer) GetDummy2Method2(ctx context.Context, req *GetDummy2Method2Request) (*Dummy2Method2Object, error) {

	if err := s.recaptchaServer.ValidateRecaptcha(ctx); err != nil {
		return nil, err
	}

	return s.Dummy2Server.GetDummy2Method2(ctx, req)
}
