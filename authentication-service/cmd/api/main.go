package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tnaucoin/zentask/authentication-service/pkg/db"
)

var (
	port = "80"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	// init db
	gormDB := db.Init()
	dbHandler := db.New(gormDB)
	svc := Init(&dbHandler)

	r := CreateRouter(svc)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run server in go routine so it's non blocking
	go func() {
		fmt.Printf("Starting server on port: %s\n", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// accept graceful shutdown via SIGINT
	signal.Notify(c, os.Interrupt)
	//Block until we recieve the signal
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	//wait for the timeout deadline
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
