package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "UsersManagement/db/sqlc"
	pbu "UsersManagement/pb/users"
	"UsersManagement/util"
)

func (s *Server) SignUp(ctx context.Context, req *pbu.SignUpRequest) (*pbu.SignUpResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Fail to Verify User")
	}

	Hash, err := util.Hash(req.GetPassword(), payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to hash password")
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
		return nil, status.Errorf(codes.Internal, "Fail to create user")
	}

	return &pbu.SignUpResponse{
		Email:       user.Email,
		FristName:   user.FristName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Msg:         "User created successfully",
	}, nil
}

func (s *Server) UserInformations(ctx context.Context, req *pbu.UserInformationRequest) (*pbu.UserInformationResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Fail to Verify User")
	}

	user, err := s.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	return &pbu.UserInformationResponse{
		Email:       user.Email,
		FristName:   user.FristName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pbu.LoginRequest) (*pbu.LoginResponse, error) {
	hashPassword, err := s.store.UserLogin(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	valid := util.VerifyHashPassword(req.GetPassword(), req.GetEmail(), hashPassword)
	if !valid {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid password")
	}

	token, payload, err := s.token.CreateToken(req.GetEmail(), "user")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Fail to Create payload")
	}

	data := db.CreateSessionParams{
		SessionID: util.CreateUUID(),
		Email:     payload.Email,
		Token:     token,
		Status:    util.GenerateSessionStatus(),
		ExpiresAt: util.GenerateDate(),
	}

	session, err := s.store.CreateSession(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to Save Session")
	}

	return &pbu.LoginResponse{
		SessionId: session.SessionID.String(),
		Token:     session.Token,
	}, nil
}

func (s *Server) ResetPasswordAfterLogin(ctx context.Context, req *pbu.ResetPasswordAfterLoginRequest) (*pbu.ResetResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	newPasswordHash, err := util.Hash(req.GetNewPassword(), payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to hash password")
	}

	data := db.ResetPasswordParams{
		Email:        payload.Email,
		PasswordHash: newPasswordHash,
	}

	err = s.store.ResetPassword(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to reset password")
	}

	return &pbu.ResetResponse{
		Success: true,
	}, nil
}

func (s *Server) ResetPhoneNumberAfterLogin(ctx context.Context, req *pbu.ResetPhoneNumberAfterLoginRequest) (*pbu.ResetResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	data := db.ResetPhoneNumberParams{
		Email:       payload.Email,
		PhoneNumber: req.GetNewPhoneNumber(),
	}

	err = s.store.ResetPhoneNumber(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to reset Phone Number")
	}

	return &pbu.ResetResponse{
		Success: true,
	}, nil
}

func (s *Server) ResetEmailAfterLogin(ctx context.Context, req *pbu.ResetEmailAfterLoginRequest) (*pbu.ResetResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	data := db.ResetEmailParams{
		Email:       payload.Email,
		PhoneNumber: req.GetNewEmail(),
	}

	err = s.store.ResetEmail(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to reset email")
	}

	return &pbu.ResetResponse{
		Success: true,
	}, nil
}
