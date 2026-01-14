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
    ON c.status IN ('new', 'active')
WHERE NOT EXISTS (
    SELECT 1
    FROM redemptions r
    WHERE r.user_id = u.id
        AND r.coupon_id = c.id
);

-- name: GetRedemptionsForUser :many
SELECT 
    u.id AS user_id,
    u.hive_id,
    u.server,
    c.id AS coupon_id,
    c.code
FROM
    users u
CROSS JOIN coupons c
WHERE u.hive_id = $1
    AND c.status IN ('new', 'active')
    AND NOT EXISTS (
        SELECT 1
        FROM redemptions r
        WHERE r.user_id = u.id
            AND r.coupon_id = c.id
    );

-- name: AddRedemptions :copyfrom
INSERT INTO redemptions(user_id, coupon_id, status, response_message)
VALUES ($1, $2, $3, $4);

