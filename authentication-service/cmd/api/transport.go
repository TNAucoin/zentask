package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type signUpRequest struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"hash"`
	FirstName    string `json"firstName"`
	LastName     string `json:"lastName"`
}

type signUpResponse struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	Err       error     `json:"error,omitempty"`
}

type signInRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"hash"`
}

type signInResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

func decodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request signInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeSignUpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(signUpResponse)
	if res.Err != nil {
		return json.NewEncoder(w).Encode(res.Err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(res)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	switch err {
	case ErrUserExists:
		w.WriteHeader(http.StatusConflict)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}
