-- name: CreateToken :one
INSERT INTO TOKENS (
    token_id,
    email,
    token,
    status,
    EXPIRES_AT
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;