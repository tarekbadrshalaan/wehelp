package ums

import (
	"encoding/json"
	"net/http"
	"time"

	bllums "wehelp-api/bll/ums"
	"wehelp-api/config"
	dtoums "wehelp-api/dto/ums"

	"wehelp-api/api/helper"
	"wehelp-api/logger"

	"github.com/julienschmidt/httprouter"
)

var validationDuration time.Duration

func init() {
	validationDuration = time.Duration(config.Configuration().ValidationDuration) * time.Minute
}

func ConfigAuthorizationRouter() []helper.Route {
	return []helper.Route{
		helper.Route{Method: "POST", Path: "/auths/jwt-token", Handle: generateJwtToken, IsLogging: true},
		helper.Route{Method: "GET", Path: "/auths/isvalid", Handle: isValidToken, IsLogging: true, IsAuthorized: true},
		helper.Route{Method: "GET", Path: "/auths/renew", Handle: renewToken, IsLogging: true, IsAuthorized: true},
	}
}

func generateJwtToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user dtoums.UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	signedToken, err := bllums.SignedJwtToken(&user, validationDuration)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusUnauthorized)
		return
	}

	helper.WriteResponseJSON(w, bllums.JwtToken{Token: signedToken}, http.StatusOK)
}

func isValidToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	userIDHeader := req.Header.Get("user-id")
	_, err := bllums.IsValid(authHeader, userIDHeader)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusUnauthorized)
		return
	}
	helper.WriteResponseJSON(w, bllums.JwtToken{Token: authHeader}, http.StatusOK)
}

func renewToken(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	userIDHeader := req.Header.Get("user-id")
	signedToken, err := bllums.SignedJwtTokenRenewed(authHeader, userIDHeader)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusUnauthorized)
		return
	}
	helper.WriteResponseJSON(w, bllums.JwtToken{Token: signedToken}, http.StatusOK)
}
