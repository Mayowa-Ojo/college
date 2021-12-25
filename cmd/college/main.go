package main

import (
	"college"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	migrateDown bool
)

func init() {
	flag.BoolVar(&migrateDown, "migrate-down", false, "specify if a down migration should run when server is started")

	flag.Parse()
}

func main() {
	context := context.Background()
	cfg := college.LoadConfig()

	client := college.NewEntClient(cfg)
	defer client.Close()

	college.DumpMigrations(context, client)

	if migrateDown {
		college.SchemaMigrateDown(context, client)
	} else {
		college.SchemaMigrateUp(context, client)
	}

	router := college.Run(client)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %s", err)
		}
	}()

	college.GracefulShutdown(srv)
}
