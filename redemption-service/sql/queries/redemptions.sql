-- name: GetUnclaimedRedemptions :many

SELECT 
    u.id,
    u.hive_id,
    u.server,
    c.id,
    c.code
FROM
    users u
JOIN coupons c
    ON c.status = 'new'
WHERE NOT EXISTS (
    SELECT 1
    FROM redemptions r
    WHERE r.user_id = u.id
        AND r.coupon_id = c.id
);

-- name: AddRedemptions :copyfrom
INSERT INTO redemptions(user_id, coupon_id, status, response_message)
VALUES ($1, $2, $3, $4);

