-- name: AddCoupon :exec
INSERT INTO Coupons (code) VALUES ($1) ON CONFLICT (code) DO NOTHING;

-- name: GetAllCoupons :many
SELECT * FROM Coupons;

-- name: GetCouponByCode :one
SELECT * FROM Coupons WHERE code = $1;

-- name: DeleteCouponById :exec
DELETE FROM Coupons WHERE id = $1;

-- name: UpdateCouponStatus :exec
UPDATE Coupons SET 
    status = $1, 
    updated_at = NOW()
WHERE
    id = $2;