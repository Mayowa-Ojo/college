package college

import (
	"college/ent"
	"college/ent/migrate"
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewEntClient(cfg *Config) *ent.Client {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	client, err := ent.Open(cfg.DBDriver, dsn)
	if err != nil {
		log.Fatalf("[ENT]: error connecting to db %s", err)
	}

	log.Println("[ENT]: connected to database")

	return client
}

func SchemaMigrateUp(ctx context.Context, client *ent.Client) {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("[ENT]: error running migration %s", err)
	}
}

func SchemaMigrateDown(ctx context.Context, client *ent.Client) {
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("[ENT]: error running migration %s", err)
	}
}

func DumpMigrations(ctx context.Context, client *ent.Client) {
	file, err := os.Create("migrations.sql")
	if err != nil {
		log.Fatalf("[ENT]: error reading/creating migrations file %s", err)
	}

	if err := client.Schema.WriteTo(ctx, file); err != nil {
		log.Fatalf("[ENT]: error writing migrations to file %s", err)
	}
}
