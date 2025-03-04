package gapi

import (
	"fmt"

	db "UsersManagement/db/sqlc"
	pbu "UsersManagement/pb/users"
	pbv "UsersManagement/pb/verification"
	"UsersManagement/token"
	"UsersManagement/util"
)

type Server struct {
	pbu.UnimplementedUserServiceServer
	pbv.UnimplementedVerificationServer
	config util.Config
	token  token.Maker
	store  db.Store
}

func ServerSetup(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.KeySeed)
	if err != nil {
		return nil, fmt.Errorf("Fail to Create Key Pair: %v", err)
	}

	return &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
	}, nil
}
