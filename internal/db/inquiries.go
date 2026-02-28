package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Inquiry represents a contact form submission.
type Inquiry struct {
	ID         int64
	Name       string
	Email      string
	Phone      string
	TripSlug   string
	TripName   string
	Dates      string
	PartySize  string
	Experience string
	Message    string
	Status     string // new, contacted, booked, archived
	Notes      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// CreateInquiry inserts a new inquiry and returns its ID.
func (s *Store) CreateInquiry(inq *Inquiry) (int64, error) {
	res, err := s.db.Exec(`
		INSERT INTO inquiries (name, email, phone, trip_slug, trip_name, dates, party_size, experience, message, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 'new')`,
		inq.Name, inq.Email, inq.Phone, inq.TripSlug, inq.TripName, inq.Dates, inq.PartySize, inq.Experience, inq.Message,
	)
	if err != nil {
		return 0, fmt.Errorf("create inquiry: %w", err)
	}
	return res.LastInsertId()
}

// GetInquiry returns a single inquiry by ID.
func (s *Store) GetInquiry(id int64) (*Inquiry, error) {
	inq := &Inquiry{}
	err := s.db.QueryRow(`
		SELECT id, name, email, phone, trip_slug, trip_name, dates, party_size, experience, message, status, notes, created_at, updated_at
		FROM inquiries WHERE id = ?`, id,
	).Scan(&inq.ID, &inq.Name, &inq.Email, &inq.Phone, &inq.TripSlug, &inq.TripName, &inq.Dates, &inq.PartySize, &inq.Experience, &inq.Message, &inq.Status, &inq.Notes, &inq.CreatedAt, &inq.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get inquiry %d: %w", id, err)
	}
	return inq, nil
}

// ListInquiries returns inquiries filtered by status (empty string = all), ordered by newest first.
func (s *Store) ListInquiries(status string) ([]Inquiry, error) {
	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = s.db.Query(`
			SELECT id, name, email, phone, trip_slug, trip_name, dates, party_size, experience, message, status, notes, created_at, updated_at
			FROM inquiries WHERE status = ? ORDER BY created_at DESC`, status)
	} else {
		rows, err = s.db.Query(`
			SELECT id, name, email, phone, trip_slug, trip_name, dates, party_size, experience, message, status, notes, created_at, updated_at
			FROM inquiries ORDER BY created_at DESC`)
	}
	if err != nil {
		return nil, fmt.Errorf("list inquiries: %w", err)
	}
	defer rows.Close()

	var inquiries []Inquiry
	for rows.Next() {
		var inq Inquiry
		if err := rows.Scan(&inq.ID, &inq.Name, &inq.Email, &inq.Phone, &inq.TripSlug, &inq.TripName, &inq.Dates, &inq.PartySize, &inq.Experience, &inq.Message, &inq.Status, &inq.Notes, &inq.CreatedAt, &inq.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan inquiry: %w", err)
		}
		inquiries = append(inquiries, inq)
	}
	return inquiries, rows.Err()
}

// UpdateInquiryStatus sets the status and updated_at for an inquiry.
func (s *Store) UpdateInquiryStatus(id int64, status string) error {
	valid := map[string]bool{"new": true, "contacted": true, "booked": true, "archived": true}
	if !valid[status] {
		return fmt.Errorf("invalid status: %s", status)
	}
	_, err := s.db.Exec(`UPDATE inquiries SET status = ?, updated_at = datetime('now') WHERE id = ?`, status, id)
	if err != nil {
		return fmt.Errorf("update inquiry status: %w", err)
	}
	return nil
}

// UpdateInquiryNotes sets the notes and updated_at for an inquiry.
func (s *Store) UpdateInquiryNotes(id int64, notes string) error {
	_, err := s.db.Exec(`UPDATE inquiries SET notes = ?, updated_at = datetime('now') WHERE id = ?`, notes, id)
	if err != nil {
		return fmt.Errorf("update inquiry notes: %w", err)
	}
	return nil
}

// CountInquiries returns the count of inquiries matching a status (empty = all).
func (s *Store) CountInquiries(status string) (int, error) {
	var count int
	var err error
	if status != "" {
		err = s.db.QueryRow("SELECT COUNT(*) FROM inquiries WHERE status = ?", status).Scan(&count)
	} else {
		err = s.db.QueryRow("SELECT COUNT(*) FROM inquiries").Scan(&count)
	}
	if err != nil {
		return 0, fmt.Errorf("count inquiries: %w", err)
	}
	return count, nil
}

// RecentInquiries returns the N most recent inquiries.
func (s *Store) RecentInquiries(limit int) ([]Inquiry, error) {
	rows, err := s.db.Query(`
		SELECT id, name, email, phone, trip_slug, trip_name, dates, party_size, experience, message, status, notes, created_at, updated_at
		FROM inquiries ORDER BY created_at DESC LIMIT ?`, limit)
	if err != nil {
		return nil, fmt.Errorf("recent inquiries: %w", err)
	}
	defer rows.Close()

	var inquiries []Inquiry
	for rows.Next() {
		var inq Inquiry
		if err := rows.Scan(&inq.ID, &inq.Name, &inq.Email, &inq.Phone, &inq.TripSlug, &inq.TripName, &inq.Dates, &inq.PartySize, &inq.Experience, &inq.Message, &inq.Status, &inq.Notes, &inq.CreatedAt, &inq.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan inquiry: %w", err)
		}
		inquiries = append(inquiries, inq)
	}
	return inquiries, rows.Err()
}
