-- name: GetUnclaimedRedemptions :many

SELECT 
    u.id,
    u.hive_id,
    u.server,
    c.id,
    c.code
FROM
    users u
CROSS JOIN
    coupons c
LEFT JOIN
    redemptions r ON u.id = r.user_id AND r.coupon_id = c.id
WHERE
    r.id IS NULL;


-- name: AddRedemptions :copyfrom
INSERT INTO redemptions(user_id, coupon_id, status, response_message)
VALUES ($1, $2, $3, $4);