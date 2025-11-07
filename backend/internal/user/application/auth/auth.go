// Package auth user token
package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(ctx context.Context, uid uint64, username string) string {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("zzHq+W8NfMpauvJJFB+ufl2guP6sgq59Ys4w66+YqiE=")
	claims := &Claims{
		UID:      uid,
		Username: username,
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
