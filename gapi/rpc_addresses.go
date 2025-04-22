package gapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	db "UsersManagement/db/sqlc"
	pba "UsersManagement/pb/addresses"
	"UsersManagement/util"
)

func (s *Server) CreateAddress(ctx context.Context, req *pba.CreateAddressRequest) (*pba.CreateAddressResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	user, err := s.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	data := db.CreateAddressParams{
		AddressID: util.CreateUUID(),
		UserID:    user.UserID,
		FlatFloor: pgtype.Text{String: req.GetFlatFloor(), Valid: true},
		Building:  pgtype.Text{String: req.GetBuilding(), Valid: true},
		Street:    req.GetStreet(),
		District:  req.GetDistrict(),
		Region:    req.GetRegion(),
		Country:   req.GetCountry(),
		ZipCode:   pgtype.Text{String: req.GetZipCode(), Valid: true},
		IsDefault: req.GetIsDefault(),
	}

	address, err := s.store.CreateAddress(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to create address")
	}

	return &pba.CreateAddressResponse{
		Address: &pba.Address{
			AddressId: address.AddressID.String(),
			FlatFloor: address.FlatFloor.String,
			Building:  address.Building.String,
			Street:    address.Street,
			District:  address.District,
			Region:    address.Region,
			Country:   address.Country,
			ZipCode:   address.ZipCode.String,
			IsDefault: address.IsDefault,
		},
	}, nil
}

func (s *Server) DeleteAddress(ctx context.Context, req *pba.DeleteAddressRequest) (*emptypb.Empty, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	user, err := s.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	addressId, err := uuid.Parse(req.GetAddressId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid address ID")
	}

	data := db.DeleteAddressParams{
		AddressID: addressId,
		UserID:    user.UserID,
	}

	err = s.store.DeleteAddress(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to delete address")
	}

	return nil, nil
}

func (s *Server) GetAddress(ctx context.Context, req *pba.GetAddressRequest) (*pba.GetAddressResponse, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	user, err := s.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	addresslist, err := s.store.GetAddressesByUserID(ctx, user.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get address list")
	}

	return &pba.GetAddressResponse{
		Address: GetAddressType(addresslist),
	}, nil
}

func (s *Server) ResetDefaultAddress(ctx context.Context, req *pba.ResetDefaultAddressRequest) (*emptypb.Empty, error) {
	payload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authorize user")
	}

	user, err := s.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get user information")
	}

	err = s.store.SetAllAddresstoFalse(ctx, user.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to set all addresses to false")
	}

	addressId, err := uuid.Parse(req.GetAddressId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid address ID")
	}

	data := db.SetDefaultAddressParams{
		UserID:    user.UserID,
		AddressID: addressId,
	}

	err = s.store.SetDefaultAddress(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to set default address")
	}

	return nil, nil
}
