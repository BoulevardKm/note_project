package handler

import (
	"github.com/gin-gonic/gin"
	jwtmanager "jwt_manager"
	"notes/internal/config"
)

type Handler struct {
	jwtManager *jwtmanager.JWTManager
	cfg        *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	jwtConfig := jwtmanager.JWTConfig{
		SecretKey:              cfg.JWTSecretKey,
		AccessTokenExpiration:  24,
		RefreshTokenExpiration: 168,
	}
	jwtmanager := jwtmanager.NewJWTManager(jwtConfig)

	return &Handler{
		cfg:        cfg,
		jwtManager: jwtmanager,
	}
}

func (h *Handler) CreateNote(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Создание заметки не реализовано",
	})
}

func (h *Handler) GetNoteByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение заметки по ID не реализовано",
	})
}

func (h *Handler) UpdateNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Обновление заметки не реализовано",
	})
}

func (h *Handler) DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Удаление заметки не реализовано",
	})
}

func (h *Handler) GetAllNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Получение всех заметок не реализовано",
	})
}

func (h *Handler) GetJWTMiddleware() gin.HandlerFunc {
	return h.jwtManager.JWTInterceptor()
}
