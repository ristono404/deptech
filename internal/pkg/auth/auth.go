package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ristono404/deptech/internal/entity/user"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var whiteListTokens = []string{}

func GenerateJWT(email string) (res user.Token, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res.AccessToken, _ = token.SignedString(jwtKey)
	res.ExpiresAt = expirationTime
	whiteListTokens = append(whiteListTokens, res.AccessToken)
	return
}
func ValidateToken(signedToken string) (err error) {

	if signedToken == "" {
		err = errors.New("access token is required")
		return
	}

	if !exists(whiteListTokens, signedToken) {
		err = errors.New("unauthorized")
		return
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		remove(whiteListTokens, signedToken)
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		remove(whiteListTokens, signedToken)
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		remove(whiteListTokens, signedToken)
		err = errors.New("token expired")
		return
	}
	return
}

func exists(s []string, r string) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func Logout(signedToken string) {
	whiteListTokens = remove(whiteListTokens, signedToken)
}
