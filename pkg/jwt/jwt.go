package jwt

import (
	"MyProgy/infrastructure/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id       int
	Username string
	Email    string
	jwt.RegisteredClaims
}

func GenerateToken(userID int, name, email string) (string, error) {
	claims := &Claims{
		Id:       userID,
		Username: name,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "MyProgy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.GetJwtEnv().JWTKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtEnv().JWTKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
