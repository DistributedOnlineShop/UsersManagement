-- name: CreateSession :one
INSERT INTO session (
    session_id,
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