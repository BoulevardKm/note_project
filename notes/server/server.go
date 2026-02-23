package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"notes/handler"
	"notes/internal/config"
	"notes/routes"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("конфигурация сервера не может быть nil")
	}
	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("не удалось создать обработчик сервера")
	}
	fmt.Println("Обработчик сервера успешно создан")

	router := routes.SetupRouter(handler)

	return &Server{
		cfg:    cfg,
		router: router,
	}, nil
}

func (s *Server) Start() error {
	fmt.Printf("Сервер запускается на %s:%s\n", s.cfg.Host, s.cfg.Port)
	return nil
}

func (s *Server) Stop() error {
	fmt.Println("Сервер остановлен")
	return nil
}

func (s *Server) Serve() error {
	if err := s.Start(); err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Сервер готов к обработке запросов на %s...\n", address)
	return s.router.Run(address)
}
