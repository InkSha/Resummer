package token

import (
	"errors"
	"maps"

	"github.com/golang-jwt/jwt/v5"
)

var ErrTokenInvalidOrExpired = errors.New("Token invalid or expired")
var ErrCannotParseClaims = errors.New("Cannot parse claims")

func GenerateToken(data map[string]any, secret string) (string, error) {
	claims := jwt.MapClaims{}

	maps.Copy(claims, data)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParserToken(token string, secret string) (map[string]any, error) {
	parse, err := jwt.Parse(
		token,
		func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
	)

	if err != nil || !parse.Valid {
		return nil, ErrTokenInvalidOrExpired
	}

	claims, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrCannotParseClaims
	}

	result := make(map[string]any)
	maps.Copy(result, claims)

	return result, nil
}
