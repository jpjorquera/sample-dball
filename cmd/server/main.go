package main

import (
	"dballz/internal/config"
	"dballz/internal/db/sqlite"
	"dballz/internal/repository"
	"dballz/internal/server"
	"dballz/internal/service"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := sqlite.New("dballz.sqlite")
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	srv := server.NewServer(cfg, userService)
	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
