package Utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var mySigningKey = []byte(os.Getenv("HGJ_JWT_SIGNING_KEY"))

type UserClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64) (string, error) {
	claims := UserClaims{id, jwt.RegisteredClaims{
		Issuer:   "HGJ",
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateTokenWithCustomClaims(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
