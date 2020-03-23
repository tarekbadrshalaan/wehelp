package api

import (
	"net/http"
	"wehelp-api/api/helper"

	apiums "wehelp-api/api/ums"

	"github.com/julienschmidt/httprouter"
)

// NewRouter :creates a new router instance and iterate through all the Routes to get each’s
// Route’s Method, Pattern and Handle and registers a new request handle.
func NewRouter() http.Handler {
	router := httprouter.New()

	for _, r := range apiums.ConfigAuthorizationRouter() {
		router.Handle(r.Method, r.Path, helper.Logmid(r.Handle))
	}

	for _, r := range apiums.ConfigUseresRouter() {
		router.Handle(r.Method, r.Path, helper.Logmid(r.Handle))
	}

	for _, r := range apiums.ConfigAddressesRouter() {
		router.Handle(r.Method, r.Path, helper.Logmid(r.Handle))
	}

	return router
}
