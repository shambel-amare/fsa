package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// this is where all dependency injections happen
func Initiate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf(`{"level":"fatal","msg":"failed to initialize sample logger: %v"}`, err)
		os.Exit(1)
	}

	logger.Info("initializing config")
	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		logger.Info(fmt.Sprintf("config name is set to %s", configName))
	} else {
		logger.Info("using default config name 'config'")
	}
	InitConfig(configName, "config", logger)
	logger.Info("config initialized")

	// Initiate the packages
	pkgLayer := InitiatePackageLayer()
	// init domain layer which have dependency of searcher pkg
	useCaseLayer := InitiateDomainLayer(pkgLayer)

	handler := InitHandlers(useCaseLayer, *logger)

	server := gin.New()

	server.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	server.Use(ginzap.RecoveryWithZap(logger, true))
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // frontend api for production
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	srv := &http.Server{
		Addr:              viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		ReadHeaderTimeout: viper.GetDuration("server.read_header_timeout"),
		Handler:           server,
	}
	apiGroup := server.Group("/api/v1")
	InitRoutes(apiGroup, handler)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	// run the server
	go func() {
		logger.Info("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}()
	sig := <-quit
	logger.Info(fmt.Sprintf("server shutting down with signal %v", sig))
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("server.timeout"))
	defer cancel()

	logger.Info("shutting down server")

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Fatal(fmt.Sprintf("error while shutting down server: %v", err))
	} else {
		logger.Info("server shutdown complete")
	}

}
