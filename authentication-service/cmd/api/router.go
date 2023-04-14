package main

import (
	"github.com/gorilla/mux"
)

func CreateRouter(svc *authenticationService) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/signup", CreateSignUpHandler(svc)).Methods("POST")
	return r
}
