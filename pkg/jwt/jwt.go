package jwt

import jwt2 "github.com/golang-jwt/jwt/v5"

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return jwt2.NewWithClaims(jwt2.SigningMethodHS512, claims).SignedString(secret)
}
