package security

import (
	"go_grab/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "khanhdanghocgolang"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//key bất kì, bảo mật
	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}
	return result, nil
}
