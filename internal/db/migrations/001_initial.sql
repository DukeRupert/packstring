-- 001_initial.sql
-- Creates core tables for inquiries, payments, and deposit configuration.

CREATE TABLE IF NOT EXISTS inquiries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL DEFAULT '',
    trip_slug TEXT NOT NULL DEFAULT '',
    trip_name TEXT NOT NULL DEFAULT '',
    dates TEXT NOT NULL DEFAULT '',
    party_size TEXT NOT NULL DEFAULT '',
    experience TEXT NOT NULL DEFAULT '',
    message TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'new' CHECK(status IN ('new','contacted','booked','archived')),
    notes TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT (datetime('now')),
    updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_inquiries_status ON inquiries(status);
CREATE INDEX IF NOT EXISTS idx_inquiries_created ON inquiries(created_at DESC);

CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    inquiry_id INTEGER NOT NULL REFERENCES inquiries(id),
    stripe_session_id TEXT NOT NULL DEFAULT '',
    stripe_payment_intent TEXT NOT NULL DEFAULT '',
    amount_cents INTEGER NOT NULL,
    currency TEXT NOT NULL DEFAULT 'usd',
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending','paid','failed','refunded')),
    customer_email TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT (datetime('now')),
    paid_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_payments_inquiry ON payments(inquiry_id);
CREATE INDEX IF NOT EXISTS idx_payments_session ON payments(stripe_session_id);

CREATE TABLE IF NOT EXISTS deposit_config (
    trip_slug TEXT PRIMARY KEY,
    trip_name TEXT NOT NULL,
    amount_cents INTEGER NOT NULL DEFAULT 0,
    enabled INTEGER NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS schema_version (
    version INTEGER PRIMARY KEY
);

INSERT INTO schema_version (version) VALUES (1);
