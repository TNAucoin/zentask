package main

import (
	"github.com/gorilla/mux"
)

func CreateRouter(svc authenticationService) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/signup", CreateSignUpHandler(svc)).Methods("POST")
	r.Handle("/signin", CreateSignInHandler(svc)).Methods("POST")
	r.Handle("/refresh", CreateRefreshTokenHandler(svc)).Methods("POST")
	return r
}
