package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jumaniyozov/grest/internal/user"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("Initializing router")
	router := httprouter.New()

	userHandler := user.NewHandler()
	userHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	port := "8080"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting application on port: ", port)
	log.Fatalln(server.Serve(listener))
}
