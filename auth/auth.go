package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint32 `json:"role"`
	jwt.StandardClaims
}

// GENERATE JWT Token
func GenerateJWT(email, username string, role uint32) (tokenString string, err error) {
	expTime := time.Now().Add(1 * time.Hour)
	// expTime := time.Now().Add(10 * time.Decond)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}

func ValidateToken(signedToken string) (email string, role uint32, err error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	// fmt.Println("token : ", token)

	claims, ok := token.Claims.(*JWTClaim)

	// fmt.Println("claims : ", claims)

	// jika claims gagal
	if !ok {
		err = errors.New("Could not pass claims from token")
		return
	}

	// jika token expired
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}

	role = claims.Role

	email = claims.Email

	return
}
