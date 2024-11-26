package app

import (
	"Songs/Song-library/config"
	"Songs/Song-library/pkg/postgres"
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/gommon/log"
)

// @title           Account Management Service
// @version         1.0
// @description     This is a service for managing accounts, reservations, products and operations.

// @contact.name   Changaz Danial
// @contact.email  changaz.d@gmail.com

// @host      localhost:8089
// @BasePath  /

// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description					JWT token

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Initializing postgres...")
	//SetLogrus(cfg.Log.Level)

	// Repositories
	//log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}
	defer pg.Close()

	// // Repositories
	// logger.Info("Initializing repositories...")
	// repositories := repo.NewRepositories(pg)

	// // Services dependencies
	// logger.Info("Initializing services...")
	// deps := service.ServicesDependencies{
	// 	Repos:    repositories,
	// 	GDrive:   gdrive.New(cfg.WebAPI.GDriveJSONFilePath),
	// 	Hasher:   hasher.NewSHA1Hasher(cfg.Hasher.Salt),
	// 	SignKey:  cfg.JWT.SignKey,
	// 	TokenTTL: cfg.JWT.TokenTTL,
	// }
	// services := service.NewServices(deps)

	// // Echo handler
	// logger.Info("Initializing handlers and routes...")
	// handler := echo.New()
	// // setup handler validator as lib validator
	// handler.Validator = validator.NewCustomValidator()
	// v1.NewRouter(handler, services)

	// // HTTP server
	// logger.Info("Starting http server...")
	// logger.Debug("Server port: %s", cfg.HTTP.Port)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// // Waiting signal
	// logger.Info("Configuring graceful shutdown...")
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// select {
	// case s := <-interrupt:
	// 	logger.Info("app - Run - signal: " + s.String())
	// case err = <-httpServer.Notify():
	// 	logger.Error("app - Run - httpServer.Notify: %w", err)
	// }

	// // Graceful shutdown
	// log.Info("Shutting down...")
	// err = httpServer.Shutdown()
	// if err != nil {
	// 	logger.Error("app - Run - httpServer.Shutdown: %w", err)
	// }
}
