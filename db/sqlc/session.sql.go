// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: session.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
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
) RETURNING session_id, email, token, status, created_at, expires_at
`

type CreateSessionParams struct {
	SessionID uuid.UUID        `json:"session_id"`
	Email     string           `json:"email"`
	Token     string           `json:"token"`
	Status    string           `json:"status"`
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.SessionID,
		arg.Email,
		arg.Token,
		arg.Status,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.SessionID,
		&i.Email,
		&i.Token,
		&i.Status,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}
