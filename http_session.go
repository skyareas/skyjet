package skyjet

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type HttpRequestSession struct {
	values  map[string]string
	isValid bool
}

func NewSession(req *http.Request) *HttpRequestSession {
	sesCfg := app.cfg.Http.Session

	ses := HttpRequestSession{map[string]string{}, false}

	c, err := req.Cookie(sesCfg.CookieName)
	if err != nil {
		return &ses
	}

	if c.Value == "" {
		return &ses
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(sesCfg.Secret), nil
	})
	if err != nil {
		return &ses
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for k, v := range claims {
			if value, valid := v.(string); valid {
				ses.values[k] = value
			}
		}
	}

	ses.isValid = true

	return &ses
}

func (s *HttpRequestSession) IsValid() bool {
	return s.isValid
}

func (s *HttpRequestSession) String() (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range s.values {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(app.cfg.Http.Session.Secret))
}

func (s *HttpRequestSession) Cookie() (*http.Cookie, error) {
	str, err := s.String()
	if err != nil {
		return nil, err
	}
	return &http.Cookie{
		Name:  app.cfg.Http.Session.CookieName,
		Value: str,
		Path:  "/",
	}, nil
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
