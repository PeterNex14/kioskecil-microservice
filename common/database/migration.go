package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
)

// RunMigrations executes pending SQL migrations using the embedded filesystem.
func RunMigrations(db *sql.DB, migrationsFS embed.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	slog.Info("Running embedded database migrations...", "dir", dir)
	if err := goose.Up(db, dir); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	slog.Info("Embedded database migrations applied successfully")
	return nil
}
