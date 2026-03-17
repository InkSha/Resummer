package auth

import (
	"time"

	"github.com/InkSha/Resummer/pkg/token"
)

func GenToken(id string, secret string) (string, error) {
	return token.GenerateToken(
		map[string]any{
			token.UserIDKey:  id,
			token.IatKey:     time.Now().Unix(),
			token.ExpiredKey: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		secret,
	)
}

func ParserToken(raw string, secret string) (map[string]any, error) {
	return token.ParserToken(raw, secret)
}
