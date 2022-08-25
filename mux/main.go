package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/camping/server"
	"github.com/gorilla/handlers"
)

func main() {

	server := server.New(server.Local(true))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, server),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("calling http.Server.Shutdown()")
		err := s.Shutdown(context.Background())
		if err != nil {
			log.Printf("http.Server.Shutdown returned err '%s'", err.Error())
		}
		log.Printf("now os.Exit()")
		os.Exit(1)
	}()
	log.Println("Starting the server end point")
	log.Fatal(s.ListenAndServe())

}
