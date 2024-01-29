package app

import (
	"context"
	"errors"
	"github.com/YANcomp/yanco-backend/internal/config"
	delivery "github.com/YANcomp/yanco-backend/internal/delivery/http"
	"github.com/YANcomp/yanco-backend/internal/server"
	"github.com/YANcomp/yanco-backend/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Creatly API
// @version 1.0
// @description REST API for Creatly App

// @host localhost:8000
// @BasePath /api/v1/
func Run(configPath string) {
	//init configs
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}

	handlers := delivery.NewHandler()

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
