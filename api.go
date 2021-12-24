package main

import (
	"context"
	"ent-demo/ent"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const TimeoutWindow = 5 * time.Second

func Run(client *ent.Client) *gin.Engine {
	router := gin.New()

	router.Use(Logger(logrus.New()), gin.Recovery())

	registerRoutes(router)

	return router
}

func GracefulShutdown(srv *http.Server) {
	close := make(chan os.Signal, 1)

	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	<-close

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), TimeoutWindow)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forcing server shutdown: %s", err)
	}

	log.Println("exiting server...")
}
