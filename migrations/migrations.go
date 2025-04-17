package migrations

import (
	"context"
	"embed"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

//go:embed schema/*.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}

func Migrate(db *bun.DB) {
	m := migrate.NewMigrator(db, Migrations)
	// Initialize migrations (create migrations table if it doesn't exist)
	if err := m.Init(context.Background()); err != nil {
		log.Fatalf("migrator.Init failed: %v", err)
	}

	group, err := m.Migrate(context.Background())
	if err != nil {
		log.Fatalf("migrator.Migrate failed: %v", err)
	}

	// Check if any migrations were applied
	if group.IsZero() {
		log.Println("no new migrations were applied")
	} else {
		log.Printf("applied migration group: %s (migrations: %v)", group, group.Migrations)
	}

	// Check applied migrations
	applied, err := m.AppliedMigrations(context.Background())
	if err != nil {
		log.Fatalf("migrator.Applied failed: %v", err)
	}
	log.Printf("total applied migrations: %d", len(applied))
}
