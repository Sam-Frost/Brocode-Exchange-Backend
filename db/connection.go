package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnectionPool(databaseUrl string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Error creating connectinn pool.. %v", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatalf("Error testing db connection... %v", err)
	}

	return pool
}
