from golang:1.20-alpine as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/authService ./cmd/api

RUN chmod +x ./bin/authService

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/bin/authService /app

CMD ["/app/authService"]

