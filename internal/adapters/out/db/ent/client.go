package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/d0lim/turnstile/ent"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewClient() (*ent.Client, error) {
	dataSourceName := os.Getenv("DB_CONNECTION_STRING")

	drv, err := sql.Open(dialect.Postgres, dataSourceName)
	if err != nil {
		return nil, err
	}
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}
