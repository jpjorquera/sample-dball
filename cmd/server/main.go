package main

import (
	"dballz/internal/config"
	"dballz/internal/model"
	"dballz/internal/repository"
	"dballz/internal/server"
	"dballz/internal/service"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := gorm.Open(sqlite.Open("data/dball.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err := dbConn.AutoMigrate(model.AllModels()...); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	charaterExternalRepository := repository.NewCharacterExternalRepository(cfg.ExternalAPIURL)
	characterDBRepository := repository.NewCharacterDBRepository(dbConn)
	characterService := service.NewCharacterService(charaterExternalRepository, characterDBRepository)

	srv := server.NewServer(cfg, characterService)
	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
