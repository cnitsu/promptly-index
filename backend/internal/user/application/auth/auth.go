// Package auth user token
package auth

import (
	"context"
	"errors"
	"time"

	"promptly-server/internal/user/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	Prompt   string `json:"prompt"`
	jwt.RegisteredClaims
}

func GenerateToken(ctx context.Context, entity *entity.UserEntity) string {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("zzHq+W8NfMpauvJJFB+ufl2guP6sgq59Ys4w66+YqiE=")
	claims := &Claims{
		UID:      entity.ID,
		Username: entity.Username,
		Prompt:   entity.Prompt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Duration(24).Hours()))),
			Issuer:    "Promptly",
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ = t.SignedString(key)
	return s
}

func ParseToken(ctx context.Context, token string) (*Claims, error) {
	tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("zzHq+W8NfMpauvJJFB+ufl2guP6sgq59Ys4w66+YqiE="), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(*Claims); ok && tk.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
