-- +goose Up

CREATE TABLE coupons (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	code TEXT UNIQUE NOT NULL, 
	status TEXT DEFAULT 'new' NOT NULL,
	first_seen_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT chk_status_type
		CHECK (status IN('new', 'active', 'expired')),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose Down

DROP TABLE coupons;