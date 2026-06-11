-- name: CreateUser :one
INSERT INTO USERS ( name, email, birth_date, password_hash, affiliate_code) VALUES ( $1, $2, $3, $4, $5) RETURNING *;

-- name: CreateUserWithReferral :one
INSERT INTO USERS ( name, email, birth_date, password_hash, affiliate_code, referrer_id) VALUES ( $1, $2, $3, $4, $5, $6) RETURNING *;


-- name: FindUserIdByAffiliateCode :one
SELECT id FROM USERS WHERE affiliate_code = $1 LIMIT 1;
