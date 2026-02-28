package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Payment represents a Stripe deposit payment.
type Payment struct {
	ID                   int64
	InquiryID            int64
	StripeSessionID      string
	StripePaymentIntent  string
	AmountCents          int
	Currency             string
	Status               string // pending, paid, failed, refunded
	CustomerEmail        string
	CreatedAt            time.Time
	PaidAt               *time.Time
}

// CreatePayment inserts a new payment record.
func (s *Store) CreatePayment(p *Payment) (int64, error) {
	res, err := s.db.Exec(`
		INSERT INTO payments (inquiry_id, stripe_session_id, stripe_payment_intent, amount_cents, currency, status, customer_email)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		p.InquiryID, p.StripeSessionID, p.StripePaymentIntent, p.AmountCents, p.Currency, p.Status, p.CustomerEmail,
	)
	if err != nil {
		return 0, fmt.Errorf("create payment: %w", err)
	}
	return res.LastInsertId()
}

// GetPaymentBySession returns a payment by its Stripe session ID.
func (s *Store) GetPaymentBySession(sessionID string) (*Payment, error) {
	p := &Payment{}
	var paidAt sql.NullTime
	err := s.db.QueryRow(`
		SELECT id, inquiry_id, stripe_session_id, stripe_payment_intent, amount_cents, currency, status, customer_email, created_at, paid_at
		FROM payments WHERE stripe_session_id = ?`, sessionID,
	).Scan(&p.ID, &p.InquiryID, &p.StripeSessionID, &p.StripePaymentIntent, &p.AmountCents, &p.Currency, &p.Status, &p.CustomerEmail, &p.CreatedAt, &paidAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get payment by session: %w", err)
	}
	if paidAt.Valid {
		p.PaidAt = &paidAt.Time
	}
	return p, nil
}

// GetPaymentsByInquiry returns all payments for a given inquiry.
func (s *Store) GetPaymentsByInquiry(inquiryID int64) ([]Payment, error) {
	rows, err := s.db.Query(`
		SELECT id, inquiry_id, stripe_session_id, stripe_payment_intent, amount_cents, currency, status, customer_email, created_at, paid_at
		FROM payments WHERE inquiry_id = ? ORDER BY created_at DESC`, inquiryID)
	if err != nil {
		return nil, fmt.Errorf("get payments by inquiry: %w", err)
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		var paidAt sql.NullTime
		if err := rows.Scan(&p.ID, &p.InquiryID, &p.StripeSessionID, &p.StripePaymentIntent, &p.AmountCents, &p.Currency, &p.Status, &p.CustomerEmail, &p.CreatedAt, &paidAt); err != nil {
			return nil, fmt.Errorf("scan payment: %w", err)
		}
		if paidAt.Valid {
			p.PaidAt = &paidAt.Time
		}
		payments = append(payments, p)
	}
	return payments, rows.Err()
}

// UpdatePaymentStatus updates the status (and optionally paid_at) for a payment.
func (s *Store) UpdatePaymentStatus(sessionID, status, paymentIntent string) error {
	valid := map[string]bool{"pending": true, "paid": true, "failed": true, "refunded": true}
	if !valid[status] {
		return fmt.Errorf("invalid payment status: %s", status)
	}

	if status == "paid" {
		_, err := s.db.Exec(`
			UPDATE payments SET status = ?, stripe_payment_intent = ?, paid_at = datetime('now')
			WHERE stripe_session_id = ?`, status, paymentIntent, sessionID)
		return err
	}

	_, err := s.db.Exec(`
		UPDATE payments SET status = ?, stripe_payment_intent = ?
		WHERE stripe_session_id = ?`, status, paymentIntent, sessionID)
	return err
}

// TotalDepositsCents returns the total amount in cents of paid deposits.
func (s *Store) TotalDepositsCents() (int64, error) {
	var total sql.NullInt64
	err := s.db.QueryRow("SELECT SUM(amount_cents) FROM payments WHERE status = 'paid'").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("total deposits: %w", err)
	}
	if total.Valid {
		return total.Int64, nil
	}
	return 0, nil
}
