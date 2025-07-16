package server

import (
	"dballz/internal/config"
	"dballz/internal/handler"
	"dballz/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func NewServer(cfg *config.Config, characterService *service.CharacterService) *Server {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	handler.RegisterHealthRoutes(r)
	handler.RegisterCharacterHandler(r, characterService)

	return &Server{
		engine: r,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(fmt.Sprintf(":%s", s.cfg.Port))
}
