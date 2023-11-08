package services

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/intwone/eda-arch-golang/internal/private_cryptography/interfaces"
)

var (
	InvalidTokenError    = errors.New("invalid token")
	TokenClaimsError     = errors.New("token claims error")
	FieldNotFoundOnToken = errors.New("field not found on token")
)

type JWTCryptography struct {
	SecretKey string
}

func NewJWTCryptography(secretKey string) interfaces.CryptographyInterface {
	return &JWTCryptography{
		SecretKey: secretKey,
	}
}

func (c *JWTCryptography) Encrypt(value string) (*string, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{"user_id": value, "exp": exp}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.SecretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (c *JWTCryptography) Decrypt(token string) (*string, error) {
	value := removeBearer(token)
	parsedToken, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(c.SecretKey), nil
		}
		return nil, InvalidTokenError
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, TokenClaimsError
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, FieldNotFoundOnToken
	}
	return &userID, nil
}

func removeBearer(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
