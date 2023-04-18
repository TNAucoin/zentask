package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeSignupEndpoint(svc authenticationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signUpRequest)
		userObject, err := svc.Signup(req.Email, req.Username, req.PasswordHash, req.FirstName, req.LastName)
		if err != nil {
			return signUpResponse{Username: req.Username, Email: req.Email, Err: err}, err
		}
		return signUpResponse{Username: req.Username, Email: req.Email, CreatedAt: userObject.CreatedAt, Err: nil}, nil

	}
}

func makeSignInEndpoint(svc authenticationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signInRequest)
		token, err := svc.Signin(req.Email, req.PasswordHash)
		if err != nil {
			return signInResponse{"", err}, err
		}
		return signInResponse{token, nil}, nil
	}
}

func makeRefreshTokenEndpoint(svc authenticationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(refreshTokenRequest)
		token, err := svc.RefreshToken(req.Jwt)
		if err != nil {
			return refreshTokenResponse{"", err}, err
		}
		return refreshTokenResponse{token, nil}, nil
	}
}
