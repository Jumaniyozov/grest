package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/jumaniyozov/grest/internal/config"
	"github.com/jumaniyozov/grest/internal/user"
	"github.com/jumaniyozov/grest/internal/user/db"
	"github.com/jumaniyozov/grest/pkg/client/mongodb"
	"github.com/jumaniyozov/grest/pkg/logging"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Creating router")
	router := httprouter.New()

	cfg := config.GetConfig()

	client, err := mongodb.NewClient(
		context.Background(),
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.Username,
		cfg.MongoDB.Password,
		cfg.MongoDB.Database,
		cfg.MongoDB.Auth_db,
	)
	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(client, cfg.MongoDB.Collection, &logger)

	users, err := storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	logger.Info(users)

	userHandler := user.NewHandler(&logger)
	userHandler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Starting server...")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("Creating socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("Listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Socket path: %s", socketPath)

	} else {
		logger.Info("Listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("Starting application on %s:%s ", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
