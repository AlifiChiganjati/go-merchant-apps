package delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlifiChiganjati/go-merchant-apps/config"
)

type Server struct {
	mux  *http.ServeMux
	host string
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		mux:  mux,
		host: host,
	}
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "We got this!")
	})
}

func (s *Server) Run() {
	s.setupRoutes()
	fmt.Println("Server is Running on port", s.host)
	if err := http.ListenAndServe(s.host, s.mux); err != nil {
		log.Fatal("server can't run")
	}
}
