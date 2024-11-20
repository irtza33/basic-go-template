package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

type GinRouter struct {
	Server *http.Server
}

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func NewGinRouter() *GinRouter {
	router := initRouter()
	return &GinRouter{
		Server: &http.Server{
			Addr:    port,
			Handler: router,
		},
	}
}

func (gr *GinRouter) Shutdown() error {
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := gr.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown abruptly: %v", err)
	}

	fmt.Println("Server shutdown gracefully")
	return nil
}