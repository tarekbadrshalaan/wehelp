package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"wehelp-api/bll"
	"wehelp-api/config"
	"wehelp-api/dto"
	"wehelp-api/logger"

	"github.com/julienschmidt/httprouter"
)

func configAuthorizationRouter() []route {
	return []route{
		route{method: "GET", path: "/auths/", handle: index},
		route{method: "POST", path: "/auths/jwt-token", handle: generateJwtToken},
		route{method: "GET", path: "/auths/isvalid", handle: isValidToken},
		route{method: "GET", path: "/auths/renew", handle: renewToken},
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Authorization-API")
}

func generateJwtToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error(err)
	}

	signedToken, err := bll.SignedJwtToken(&user, time.Duration(config.Configuration().ValidationDuration)*time.Minute)

	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	writeResponseJSON(w, bll.JwtToken{Token: signedToken}, http.StatusOK)
}

func isValidToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	_, err := bll.IsValid(authHeader)

	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(200)
}

func renewToken(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	authHeader := req.Header.Get("Authorization")
	signedToken, err := bll.SignedJwtTokenRenewed(authHeader)

	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusUnauthorized)
		return
	}
	writeResponseJSON(w, bll.JwtToken{Token: signedToken}, http.StatusOK)
}
