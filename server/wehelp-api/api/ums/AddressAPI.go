package ums

import (
	"fmt"
	"net/http"

	"wehelp-api/api/helper"
	bllums "wehelp-api/bll/ums"
	dtoums "wehelp-api/dto/ums"

	// helper "wehelp-api/api/helper"
	"wehelp-api/logger"
	"wehelp-api/utils"

	"github.com/julienschmidt/httprouter"
)

func ConfigAddressesRouter() []helper.Route {
	return []helper.Route{
		helper.Route{Method: "GET", Path: "/addresses", Handle: getAllAddresses},
		helper.Route{Method: "POST", Path: "/addresses", Handle: postAddresses},
		helper.Route{Method: "PUT", Path: "/addresses", Handle: putAddresses},
		helper.Route{Method: "GET", Path: "/addresses/:id", Handle: getAddresses},
		helper.Route{Method: "DELETE", Path: "/addresses/:id", Handle: deleteAddresses},
	}
}

func getAllAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	addresses, err := bllums.GetAllAddresses()
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteResponseJSON(w, addresses, http.StatusOK)
}

func getAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	requestID := ps.ByName("id")
	id, err := utils.ConvertID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusBadRequest)
		return
	}

	address, err := bllums.GetAddress(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find address (%v); err (%v)", id, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return
	}
	helper.WriteResponseJSON(w, address, http.StatusOK)
}

func postAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dtoums.AddressDTO{}
	if err := helper.ReadJSON(r, address); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bllums.CreateAddress(address)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteResponseJSON(w, result, http.StatusCreated)
}

func putAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dtoums.AddressDTO{}
	if err := helper.ReadJSON(r, address); err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bllums.UpdateAddress(address)
	if err != nil {
		logger.Error(err)
		helper.WriteResponseError(w, err, http.StatusBadRequest)
		return
	}
	helper.WriteResponseJSON(w, result, http.StatusOK)
}

func deleteAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	requestID := ps.ByName("id")
	id, err := utils.ConvertID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusBadRequest)
		return
	}

	err = bllums.DeleteAddress(id)
	if err != nil {
		msg := fmt.Errorf("Address with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		helper.WriteResponseError(w, msg, http.StatusNotFound)
		return

	}
	helper.WriteResponseJSON(w, true, http.StatusOK)
}
