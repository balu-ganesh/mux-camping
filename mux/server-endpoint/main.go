package main

import (
	"context"
	"flag"
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
	// log.Printf("Build date: %v Build hash: %v Build number: %v", version.Date, version.Hash, version.Build)
	// log.Printf("Starting up adapter endpoint...\n\n")

	var isLocal bool
	flag.BoolVar(&isLocal, "local", false, "set local to true if the adapter endpoint is a dev environment")
	flag.Parse()
	server := server.New(server.Local(isLocal))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers.CombinedLoggingHandler(os.Stdout, server),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("we get signal, main screen turn on! calling http.Server.Shutdown()")
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
