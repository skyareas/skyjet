package server

import (
	"fmt"
	"net/http"

	"github.com/akaahmedkamal/go-server/config"
	"github.com/golang-jwt/jwt"
)

type HttpRequestSession struct {
	values map[string]string
}

func NewSession(req *http.Request) (*HttpRequestSession, error) {
	cfg := config.Shared().Http.Session

	ses := HttpRequestSession{map[string]string{}}

	c, err := req.Cookie(cfg.CookieName)
	if err != nil {
		return &ses, nil
	}

	if c.Value == "" {
		return &ses, nil
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return &ses, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for k, v := range claims {
			if value, valid := v.(string); valid {
				ses.values[k] = value
			}
		}
	}

	return &ses, nil
}

func (s *HttpRequestSession) String() (string, error) {
	cfg := config.Shared().Http.Session
	claims := jwt.MapClaims{}
	for k, v := range s.values {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.Secret)
}

func (s *HttpRequestSession) Get(name string) string {
	return s.values[name]
}

func (s *HttpRequestSession) Lookup(name string) (string, bool) {
	value, exists := s.values[name]
	return value, exists
}

func (s *HttpRequestSession) Set(name string, value string) {
	s.values[name] = value
}
