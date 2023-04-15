module github.com/tnaucoin/zentask/authentication-service

go 1.20

require (
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0-rc.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)

replace github.com/tnaucoin/zentask/authentication-service/models => ./models

replace github.com/tnaucoin/zentask/authentication-service/pkg/token => ./pkg/token
