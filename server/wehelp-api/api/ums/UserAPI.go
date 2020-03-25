package ums

import (
	"errors"
	"fmt"
	"net/http"

	"wehelp-api/api/helper"
	bllums "wehelp-api/bll/ums"
	dtoums "wehelp-api/dto/ums"
	"wehelp-api/logger"
	"wehelp-api/utils"

	"github.com/julienschmidt/httprouter"
)

func ConfigUseresRouter() []helper.Route {
	return []helper.Route{
		// helper.Route{Method: "GET", Path: "/useres", Handle: getAllUseres},
		helper.Route{Method: "POST", Path: "/useres", Handle: postUseres, IsLogging: true},
		// helper.Route{Method: "PUT", Path: "/useres", Handle: putUseres},
		helper.Route{Method: "GET", Path: "/useres/", Handle: getUseres, IsLogging: true, IsAuthorized: true},
		// helper.Route{Method: "DELETE", Path: "/useres/:id", Handle: deleteUseres},
	}
}

func getAllUseres(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	useres, err := bllums.GetAllUseres()
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteResponseJSON(w, useres, http.StatusOK)
}

func getUseres(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userIDHeader := r.Header.Get("user-id")
	id, err := utils.ConvertID(userIDHeader)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", userIDHeader, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusBadRequest)
		return
	}

	user, err := bllums.GetUser(id)
	if err != nil {
		if errors.Is(err, utils.ErrRecordNotFound) {
			helper.WriteResponseError(w, err, http.StatusNotFound)
			return
		}
		msg := fmt.Errorf("Can’t find user (%v); err (%v)", id, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusInternalServerError)
		return
	}
	helper.WriteResponseJSON(w, user, http.StatusOK)
}

func postUseres(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := &dtoums.UserDTO{}
	if err := helper.ReadJSON(r, user); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bllums.CreateUser(user)
	if err != nil {
		if errors.Is(err, utils.ErrAlreadyExists) {
			helper.WriteResponseError(w, err, http.StatusBadRequest)
			return
		}
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteResponseJSON(w, result, http.StatusCreated)
}

func putUseres(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := &dtoums.UserDTO{}
	if err := helper.ReadJSON(r, user); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bllums.UpdateUser(user)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}
	helper.WriteResponseJSON(w, result, http.StatusOK)
}

func deleteUseres(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	requestID := ps.ByName("id")
	id, err := utils.ConvertID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusBadRequest)
		return
	}

	err = bllums.DeleteUser(id)
	if err != nil {
		msg := fmt.Errorf("User with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return

	}
	helper.WriteResponseJSON(w, true, http.StatusOK)
}
