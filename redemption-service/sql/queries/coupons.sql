-- name: AddCoupon :exec
INSERT INTO Coupons (code) VALUES ($1) ON CONFLICT (code) DO NOTHING;

-- name: GetAllCoupons :many
SELECT * FROM Coupons;

-- name: GetCouponByCode :one
SELECT * FROM Coupons WHERE code = $1;

-- name: DeleteCouponById :exec
DELETE FROM Coupons WHERE id = $1;

-- name: DeleteExpiredCoupons :exec
DELETE FROM Coupons WHERE status = 'expired';

-- name: SetCouponActive :exec
UPDATE Coupons SET 
    status = 'active', 
    updated_at = NOW()
WHERE
    id = $1;

-- name: SetCouponExpired :exec
UPDATE Coupons SET 
    status = 'expired', 
    updated_at = NOW()
WHERE
    id = $1;
