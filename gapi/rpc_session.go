package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	db "UsersManagement/db/sqlc"
	pbs "UsersManagement/pb/session"
	"UsersManagement/util"
)

func (s *Server) CreateSessionId(ctx context.Context, req *pbs.CreateSessionIdRequest) (*pbs.CreateSessionIdResponse, error) {
	token, payload, err := s.token.CreateToken(req.GetEmail(), "user")
	if err != nil {
		return nil, fmt.Errorf("Fail to Create payload: %v", err)
	}

	data := db.CreateSessionParams{
		SessionID: util.CreateUUID(),
		Email:     payload.Email,
		Token:     token,
		Status:    req.GetStatus(),
		ExpiresAt: pgtype.Timestamp{Time: payload.ExpiredAt, Valid: true},
	}

	session, err := s.store.CreateSession(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("Fail to Create Session: %v", err)
	}

	return &pbs.CreateSessionIdResponse{
		Token: session.Token,
	}, nil
}
