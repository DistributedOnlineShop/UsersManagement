package gapi

import (
	"context"
	"fmt"

	db "UsersManagement/db/sqlc"
	pbu "UsersManagement/pb/users"
	"UsersManagement/util"
)

func (s *Server) SignUp(ctx context.Context, req *pbu.SignUpRequest) (*pbu.SignUpResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("Fail to Verify User: %v", err)
	}

	Hash, err := util.Hash(req.GetPassword(), payload.Email)
	if err != nil {
		return nil, fmt.Errorf("Fail to hash password: %v", err)
	}

	data := db.CreateUserParams{
		UserID:       util.CreateUUID(),
		FristName:    req.GetFristName(),
		LastName:     req.GetLastName(),
		Email:        payload.Email,
		PhoneNumber:  req.GetPhoneNumber(),
		PasswordHash: Hash,
		Status:       "Active",
	}
	user, err := s.store.CreateUser(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("Fail to create user: %v", err)
	}

	return &pbu.SignUpResponse{
		Email:       user.Email,
		FristName:   user.FristName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Msg:         "User created successfully",
	}, nil
}
