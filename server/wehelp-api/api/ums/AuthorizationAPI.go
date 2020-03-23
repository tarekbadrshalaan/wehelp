package ums

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"wehelp-api/bll/ums"
	"wehelp-api/config"
	dtoums "wehelp-api/dto/ums"

	"wehelp-api/api/helper"
	"wehelp-api/logger"

	"github.com/julienschmidt/httprouter"
)

func ConfigAuthorizationRouter() []helper.Route {
	return []helper.Route{
		helper.Route{Method: "GET", Path: "/auths/", Handle: index},
		helper.Route{Method: "POST", Path: "/auths/jwt-token", Handle: generateJwtToken},
		helper.Route{Method: "GET", Path: "/auths/isvalid", Handle: isValidToken},
		helper.Route{Method: "GET", Path: "/auths/renew", Handle: renewToken},
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Authorization-API")
}

func generateJwtToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user dtoums.UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error(err)
	}

	signedToken, err := ums.SignedJwtToken(&user, time.Duration(config.Configuration().ValidationDuration)*time.Minute)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	helper.WriteResponseJSON(w, ums.JwtToken{Token: signedToken}, http.StatusOK)
}

func isValidToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	_, err := ums.IsValid(authHeader)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(200)
}

func renewToken(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	signedToken, err := ums.SignedJwtTokenRenewed(authHeader)

	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusUnauthorized)
		return
	}
	helper.WriteResponseJSON(w, ums.JwtToken{Token: signedToken}, http.StatusOK)
}
