package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jumaniyozov/grest/internal/user"
	"github.com/jumaniyozov/grest/pkg/logging"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Creating router")
	router := httprouter.New()

	userHandler := user.NewHandler(logger)
	userHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Starting server...")

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

	logger.Info("Starting application on port: ", port)
	logger.Fatal(server.Serve(listener))
}
