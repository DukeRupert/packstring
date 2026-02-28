package db

import (
	"fmt"
)

// DepositConfig holds per-trip deposit settings.
type DepositConfig struct {
	TripSlug    string
	TripName    string
	AmountCents int
	Enabled     bool
}

// GetDepositConfig returns the deposit config for a specific trip.
func (s *Store) GetDepositConfig(tripSlug string) (*DepositConfig, error) {
	dc := &DepositConfig{}
	var enabled int
	err := s.db.QueryRow(`
		SELECT trip_slug, trip_name, amount_cents, enabled
		FROM deposit_config WHERE trip_slug = ?`, tripSlug,
	).Scan(&dc.TripSlug, &dc.TripName, &dc.AmountCents, &enabled)
	if err != nil {
		return nil, nil // not configured yet
	}
	dc.Enabled = enabled == 1
	return dc, nil
}

// ListDepositConfigs returns all deposit configurations.
func (s *Store) ListDepositConfigs() ([]DepositConfig, error) {
	rows, err := s.db.Query(`SELECT trip_slug, trip_name, amount_cents, enabled FROM deposit_config ORDER BY trip_slug`)
	if err != nil {
		return nil, fmt.Errorf("list deposit configs: %w", err)
	}
	defer rows.Close()

	var configs []DepositConfig
	for rows.Next() {
		var dc DepositConfig
		var enabled int
		if err := rows.Scan(&dc.TripSlug, &dc.TripName, &dc.AmountCents, &enabled); err != nil {
			return nil, fmt.Errorf("scan deposit config: %w", err)
		}
		dc.Enabled = enabled == 1
		configs = append(configs, dc)
	}
	return configs, rows.Err()
}

// SaveDepositConfig upserts a deposit configuration for a trip.
func (s *Store) SaveDepositConfig(dc *DepositConfig) error {
	enabled := 0
	if dc.Enabled {
		enabled = 1
	}
	_, err := s.db.Exec(`
		INSERT INTO deposit_config (trip_slug, trip_name, amount_cents, enabled, updated_at)
		VALUES (?, ?, ?, ?, datetime('now'))
		ON CONFLICT(trip_slug) DO UPDATE SET
			trip_name = excluded.trip_name,
			amount_cents = excluded.amount_cents,
			enabled = excluded.enabled,
			updated_at = datetime('now')`,
		dc.TripSlug, dc.TripName, dc.AmountCents, enabled,
	)
	if err != nil {
		return fmt.Errorf("save deposit config: %w", err)
	}
	return nil
}
