package ums

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"wehelp-api/config"
	dtoums "wehelp-api/dto/ums"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtToken :
type JwtToken struct {
	Token string `json:"token"`
}

var signedKey = []byte(config.Configuration().SigendKey)

// SignedJwtToken return sigened auth header
func SignedJwtToken(user *dtoums.UserLoginDTO, validationDuration time.Duration) (string, error) {
	if !(user.Email == "1" && user.Password == "1") { // TODO replace with database check
		return "", errors.New("Not a registered user")
	}

	// Create a new token object, specifying signing method and the claims to be contained
	expirationTime := time.Now().Add(validationDuration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   expirationTime.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := token.SignedString(signedKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// SignedJwtTokenRenewed returns a new signed token based on old token
func SignedJwtTokenRenewed(authHeader string) (string, error) {
	token, err := IsValid(authHeader)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(time.Duration(1) * time.Minute)
	// Override token's expiration time
	token.Claims.(jwt.MapClaims)["exp"] = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token.Claims)

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := newToken.SignedString(signedKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// IsValid checks authentication header, returns its token metadata
func IsValid(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, errors.New("An Authorization header is required")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return nil, errors.New("Invalid bearerToken length")
	}

	token, err := jwt.Parse(bearerToken[1], lookupValidatingKey)
	// catch token errors
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid authentication token")
	}

	return token, nil

}

func lookupValidatingKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("There was an error")
	}
	return signedKey, nil
}
