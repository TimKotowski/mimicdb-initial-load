package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/TimKotowski/postgres-snapshotting-temporal/migrations"
)

func main() {
	config, err := pgxpool.ParseConfig("postgres://postgres_source:12345@tempy:5432/postgres_source?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to parse pgxpool config: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	db := bun.NewDB(sqldb, pgdialect.New())

	// Add query logging for debugging
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	migrations.Migrate(db)
	fmt.Println(runtime.Caller(0))
}
