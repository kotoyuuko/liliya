package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server contains config of http server
type Server struct {
	Engine       *gin.Engine
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

// Serve starts http server
func (s *Server) Serve() {
	server := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Engine,
		ReadTimeout:    time.Duration(s.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(s.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
