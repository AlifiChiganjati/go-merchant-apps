package delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/config"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/delivery/controller"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/di"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine     *gin.Engine
	host       string
	uc         di.UsecaseDI
	jwtService *jwttoken.JWTService
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	jwtService := jwttoken.NewJWTService(
		cfg.TokenConfig.IssuerName,
		cfg.TokenConfig.JwtSignatureKey,
	)
	infra, err := di.NewInfraDI(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := di.NewRepoDI(infra)
	uc := di.NewUseCaseDI(
		repo,
		jwtService,
		cfg.TokenConfig.JwtLifeTime,
	)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.APIPort)

	return &Server{
		engine:     engine,
		host:       host,
		uc:         uc,
		jwtService: jwtService,
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
	controller.NewMerchantController(s.uc.MerchantUsecase(), rg, s.jwtService).Route()
	controller.NewProductController(s.uc.ProductUsecase(), rg, s.jwtService).Route()
	controller.NewCartController(s.uc.CartUsecase(), rg, s.jwtService).Route()
	controller.NewOrderController(s.uc.OrderUsecase(), rg, s.jwtService).Route()
}

func (s *Server) Run() {
	s.setupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
