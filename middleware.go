package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func SetGoogleCaptchaResponseInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if captcha := r.Header.Get("g-recaptcha-response"); captcha != "" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "g-recaptcha-response", captcha)))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "g-recaptcha-response", generateFakeCaptchaToken())))
	})
}

func generateFakeCaptchaToken() string {
	// reCAPTCHA tokens are usually ~1000 characters
	// Let's generate 750 bytes, which base64-encodes to ~1000 chars
	buf := make([]byte, 750)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	token := base64.RawURLEncoding.EncodeToString(buf)
	return token
}
