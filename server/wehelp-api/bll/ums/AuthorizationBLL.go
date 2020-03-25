package ums

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"wehelp-api/config"
	dalums "wehelp-api/dal/ums"
	dtoums "wehelp-api/dto/ums"
	"wehelp-api/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtToken :
type JwtToken struct {
	Token string `json:"token"`
}

var signedKey = []byte(config.Configuration().SigendKey)

func checkUserEmailPassword(user *dtoums.UserLoginDTO) (*dtoums.UserDTO, bool) {
	passmd5 := utils.StringToMd5(user.Password)
	dbuser, err := dalums.GetUserByEmailPassword(user.Email, passmd5)
	if dbuser == nil || err != nil {
		return nil, false
	}
	return &dtoums.UserDTO{ID: dbuser.ID, Email: dbuser.Email, Name: dbuser.Name}, true
}

// SignedJwtToken return sigened auth header
func SignedJwtToken(user *dtoums.UserLoginDTO, validationDuration time.Duration) (string, error) {
	userDTO, valid := checkUserEmailPassword(user)
	if !valid {
		return "", errors.New("email or password is incorrect")
	}

	// Create a new token object, specifying signing method and the claims to be contained
	expirationTime := time.Now().Add(validationDuration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userDTO.ID,
		"email": userDTO.Email,
		"name":  userDTO.Name,
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
func SignedJwtTokenRenewed(authHeader string, userID string) (string, error) {
	token, err := IsValid(authHeader, userID)
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

// IsValid checks authentication header and user-id, returns its token metadata
func IsValid(authHeader string, userID string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, errors.New("An Authorization header is required")
	}

	if userID == "" {
		return nil, errors.New("An user-id header is required")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return nil, errors.New("Invalid bearer token length")
	}

	token, err := jwt.Parse(bearerToken[1], lookupValidatingKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid authentication token")
	}

	mapClaims, valid := token.Claims.(jwt.MapClaims)
	if !valid {
		return nil, errors.New("Invalid authentication token")
	}

	if fmt.Sprint(mapClaims["id"]) != userID {
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
