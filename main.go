package main

import (
	"context"
	"errors"
	"log"
	"net/http"
)

func main() {
	context := context.Background()
	cfg := LoadConfig()

	client := NewEntClient(cfg)
	defer client.Close()

	SchemaMigrateUp(context, client)

	router := Run(client)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %s", err)
		}
	}()

	GracefulShutdown(srv)
}
