package db

import "embed"

// MigrationsFS embeds all SQL migration files inside migrations/
//
//go:embed migrations/*.sql
var MigrationsFS embed.FS
