package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var QueriesInstance *Queries // Global Queries instance
var DBPool *pgxpool.Pool     // Global pgx pool

func InitDatabase() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("❌ Environment variable POSTGRES_DSN is not set")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ Could not create connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("❌ Could not connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL via pgx/v5")

	DBPool = pool
	QueriesInstance = New(pool)

}
