package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

var options = []httptransport.ServerOption{
	httptransport.ServerErrorEncoder(encodeError),
}

func CreateSignUpHandler(svc *authenticationService) *httptransport.Server {
	return httptransport.NewServer(
		makeSignupEndpoint(svc),
		decodeSignUpRequest,
		encodeSignUpResponse,
		options...,
	)
}

func CreateSignInHandler(svc *authenticationService) *httptransport.Server {
	return httptransport.NewServer(
		makeSignInEndpoint(svc),
		decodeSignInRequest,
		encodeResponse,
		options...,
	)
}
