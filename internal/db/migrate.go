package db

import (
	"embed"
	"fmt"
	"log"
)

//go:embed migrations/*.sql
var migrations embed.FS

// migrate runs any pending SQL migrations.
func (s *Store) migrate() error {
	// Check current schema version
	current := 0
	row := s.db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='schema_version'")
	var name string
	if err := row.Scan(&name); err == nil {
		// Table exists, read version
		row := s.db.QueryRow("SELECT COALESCE(MAX(version), 0) FROM schema_version")
		if err := row.Scan(&current); err != nil {
			return fmt.Errorf("read schema version: %w", err)
		}
	}

	// Run migrations in order
	needed := []struct {
		version int
		file    string
	}{
		{1, "migrations/001_initial.sql"},
	}

	for _, m := range needed {
		if m.version <= current {
			continue
		}
		data, err := migrations.ReadFile(m.file)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", m.file, err)
		}
		if _, err := s.db.Exec(string(data)); err != nil {
			return fmt.Errorf("run migration %s: %w", m.file, err)
		}
		log.Printf("[db] applied migration %s", m.file)
	}

	return nil
}
