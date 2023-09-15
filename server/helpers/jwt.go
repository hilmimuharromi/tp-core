package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"tp-core/server/config"
)

type JwtCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func CreateJwtToken(userId string) (string, error) {
	config.LoadEnv()
	secretKey := os.Getenv("SECRET_KEY")
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// we hash the jwt claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
