package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")
	return dbpool
}
