package main

import (
	"calculator/internal/app/calculator/grpc"
	"calculator/internal/app/calculator/rest"
	"calculator/internal/pkg/service"
	grpcserver "calculator/internal/server/grps"
	restserver "calculator/internal/server/rest"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Calculator API
// @version 1.0
// @description Пример REST API для калькулятора

// @host localhost:8080
// @BasePath /
func main() {
	ctx := context.Background()
	err := bootstrap(ctx)
	if err != nil {
		log.Fatalf("[main] bootstrap: %v", err)
	}
}

func bootstrap(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	logger.Info("[bootstrap] инициализация базового сервиса.")
	commonService := service.NewService(logger)

	logger.Info("[bootstrap] инициализация REST сервиса.")
	restServer := restserver.NewServerREST()
	restService := rest.NewRESTService(logger, commonService)
	restHandler := rest.NewHandler(restService)

	logger.Info("[bootstrap] инициализация gRPC сервиса.")
	grpcServer := grpcserver.NewServerGRPC()
	grpcService := grpc.NewGRPSService(logger, commonService)
	grpcserver.Registration(grpcServer.Server, grpcService)

	go func() {
		logger.Info("Старт REST сервера на порту :8080")
		if err := restServer.Run("8080", restHandler.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.WithError(err).Fatal("REST сервис закончил работу с ошибкой")
		}
	}()

	go func() {
		logger.Info("Старт gRPC сервера на порту :8090")
		if err := grpcServer.Run("8090"); err != nil {
			logger.WithError(err).Fatal("gRPC сервер закончил работу с ошибкой")
		}
	}()

	<-ctx.Done()

	logger.Info("[bootstrap] Мягкое отключение серверов...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := restServer.ShutDown(shutdownCtx); err != nil {
		logger.WithError(err).Error("REST server shutdown error")
	}

	if err := grpcServer.ShutDown(shutdownCtx); err != nil {
		logger.WithError(err).Error("gRPC server shutdown error")
	}

	logger.Info("[bootstrap] Сервер мягко закончил работу.")

	return nil
}
