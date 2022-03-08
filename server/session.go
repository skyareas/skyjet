package server

import (
	"fmt"
	"net/http"

	"github.com/akaahmedkamal/go-server/config"
	"github.com/golang-jwt/jwt"
)

type HttpRequestSession struct {
	cfg    config.HttpSessionConfig
	values map[string]string
}

func NewSession(req *http.Request) (*HttpRequestSession, error) {
	cfg := config.Shared().Http.Session

	c, err := req.Cookie(cfg.CookieName)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	ses := HttpRequestSession{cfg, map[string]string{}}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for k, v := range claims {
			if value, valid := v.(string); valid {
				ses.values[k] = value
			}
		}
	} else {
		return nil, err
	}

	return &ses, nil
}

func (s *HttpRequestSession) String() (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range s.values {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.cfg.Secret)
}
