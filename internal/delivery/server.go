package delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/config"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/delivery/controller"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/di"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
	uc     di.UsecaseDI
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := di.NewInfraDI(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := di.NewRepoDI(infra)
	uc := di.NewUseCaseDI(repo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		engine: engine,
		host:   host,
		uc:     uc,
	}
}

func (s *Server) setupRoutes() {
	rg := s.engine.Group("/api/v1")
	rg.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  http.StatusOK,
			"message": "server is running",
		})
	})

	controller.NewAuthController(s.uc.AuthUsecase(), rg).Route()
	controller.NewMerchantController(s.uc.MerchantUsecase(), rg).Route()
	controller.NewProductController(s.uc.ProductUsecase(), rg).Route()
}

func (s *Server) Run() {
	s.setupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
