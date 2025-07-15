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

	charaterExternalRepository := repository.NewCharacterExternalRepository(cfg.ExternalAPIURL)
	characterDBRepository := repository.NewCharacterDBRepository(dbConn)
	characterService := service.NewCharacterService(charaterExternalRepository, characterDBRepository)

	srv := server.NewServer(cfg, characterService)
	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
