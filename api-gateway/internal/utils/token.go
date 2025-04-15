package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenUtil interface {
	GenerateToken(userID int) (string, error)
	VerifyToken(tokenString string) (*Claims, error)
}

type JWTTokenUtil struct {
	secretKey []byte
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func NewJWTTokenUtil(secretKey string) *JWTTokenUtil {
	return &JWTTokenUtil{
		secretKey: []byte(secretKey),
	}
}

func (t *JWTTokenUtil) GenerateToken(userID int) (string, error) {
	if userID <= 0 {
		return "", errors.New("invalid user ID")
	}

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.secretKey)
}

func (t *JWTTokenUtil) VerifyToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New("token string is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return t.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
