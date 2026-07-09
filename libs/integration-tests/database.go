package integrationtests

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/libs/db"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func StartPostgres(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	ctx := context.Background()

	postgresContainer, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("start postgres container: %v", err)
	}
	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate postgres container: %v", err)
		}
	})

	dsn, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("postgres connection string: %v", err)
	}

	db, err := db.NewConnection(dsn)
	if err != nil {
		t.Fatalf("connect postgres: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("close db: %v", err)
		}
	})

	if err := db.Ping(); err != nil {
		t.Fatalf("ping postgres: %v", err)
	}
	if _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		t.Fatalf("create uuid extension: %v", err)
	}

	runMigrations(t, db)

	return db, func() {
		db.Close()
	}
}

func runMigrations(t *testing.T, db *sql.DB) {
	dir := findMigrationsDir(t)
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read migrations dir: %v", err)
	}

	var files []string
	for _, f := range entries {
		if f.IsDir() || filepath.Ext(f.Name()) != ".sql" {
			continue
		}
		name := f.Name()
		if len(name) < 7 || name[len(name)-7:] != ".up.sql" {
			continue
		}
		files = append(files, name)
	}
	sort.Strings(files)

	for _, f := range files {
		content, err := os.ReadFile(filepath.Join(dir, f))
		if err != nil {
			t.Fatalf("read migration %s: %v", f, err)
		}

		raw := string(content)
		for stmt := range strings.SplitSeq(raw, ";") {
			stmt = removeComments(stmt)
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			if _, err := db.Exec(stmt); err != nil {
				t.Fatalf("migration %s failed: %v\nSQL: %s", f, err, stmt)
			}
		}
	}
}

func removeComments(s string) string {
	var out []string
	for line := range strings.SplitSeq(s, "\n") {
		tr := strings.TrimSpace(line)
		if strings.HasPrefix(tr, "--") {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func findMigrationsDir(t *testing.T) string {
	for _, rel := range []string{"../../migrations", "../migrations", "migrations"} {
		abs, err := filepath.Abs(rel)
		if err != nil {
			continue
		}
		if info, err := os.Stat(abs); err == nil && info.IsDir() {
			return abs
		}
	}
	t.Fatal("migrations directory not found in any known location")
	return ""
}
