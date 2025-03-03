package gapi

import (
	db "UsersManagement/db/sqlc"
	pbu "UsersManagement/pb/users"
	pbv "UsersManagement/pb/verification"
	"UsersManagement/util"
)

type Server struct {
	pbu.UnimplementedUserServiceServer
	pbv.UnimplementedVerificationServer
	config util.Config
	store  db.Store
}

func ServerSetup(config util.Config, store db.Store) (*Server, error) {
	return &Server{
		config: config,
		store:  store,
	}, nil
}
