package recaptcha_module

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ReCaptchaVerifyUrl = "https://www.google.com/recaptcha/api/siteverify"
	ErrorInternal      = "Something Went Wrong !!!"
)

type ReCaptchaServer struct {
	logger             *zap.Logger
	ReCaptchaSkip      bool
	ReCaptchaSecretKey string
}

type RecaptchaResponse struct {
	Success    bool     `json:"success"`
	Score      float32  `json:"score"`
	Action     string   `json:"action"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

func (c *ReCaptchaServer) ValidateRecaptcha(ctx context.Context) error {
	if !c.ReCaptchaSkip {
		if c.ReCaptchaSecretKey == "" {
			c.logger.Error("Error in CheckValidateRecaptcha - Unable to get Recaptcha Secret Key from ENV")
			return status.Error(codes.Internal, ErrorInternal)
		}

		requestReCaptchaObjectString := ctx.Value("g-recaptcha-response").(string)
		if requestReCaptchaObjectString == "" {
			c.logger.Error("Error in CheckValidateRecaptcha - Unable to get 'g-recaptcha-response' key from context")
			return status.Error(codes.Internal, ErrorInternal)
		}

		resp, err := http.PostForm(ReCaptchaVerifyUrl, url.Values{
			"secret":   {c.ReCaptchaSecretKey},
			"response": {requestReCaptchaObjectString},
		})
		if err != nil {
			c.logger.Error("Error in CheckValidateRecaptcha - Post Form", zap.Error(err))
			return status.Error(codes.Internal, ErrorInternal)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.logger.Error("Error in CheckValidateRecaptcha - Read body", zap.Error(err))
			return status.Error(codes.Internal, ErrorInternal)
		}

		recaptchaRes := &RecaptchaResponse{}
		err = json.Unmarshal(body, recaptchaRes)
		if err != nil {
			c.logger.Error("Error in CheckValidateRecaptcha - Unmarshal JSON Recaptcha Response", zap.Error(err))
			return status.Error(codes.Internal, ErrorInternal)
		}

		if !recaptchaRes.Success {
			c.logger.Error("Error in CheckValidateRecaptcha - Captcha Validation Failed",
				zap.Strings("error-codes", recaptchaRes.ErrorCodes))
			return status.Error(codes.PermissionDenied, "ReCaptcha Validation Failed")
		}
	}
	return nil
}
