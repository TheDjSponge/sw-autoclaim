-- +goose Up
CREATE TABLE redemptions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID references users(id) ON DELETE CASCADE NOT NULL,
    coupon_id UUID references coupons(id) ON DELETE CASCADE NOT NULL,
    status TEXT NOT NULL,
    response_message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_status CHECK (status IN('success', 'failed', 'already_claimed'))
);

-- +goose Down

DROP TABLE redemptions;