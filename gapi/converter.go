package gapi

import (
	db "UsersManagement/db/sqlc"
	pba "UsersManagement/pb/addresses"
)

func GetAddressType(add []db.Address) []*pba.Address {
	var addressPb []*pba.Address
	for i, v := range add {
		addressPb[i].AddressId = v.AddressID.String()
		addressPb[i].UserId = v.UserID.String()
		addressPb[i].FlatFloor = v.FlatFloor.String
		addressPb[i].Building = v.Building.String
		addressPb[i].Street = v.Street
		addressPb[i].District = v.District
		addressPb[i].Region = v.Region
		addressPb[i].Country = v.Country
		addressPb[i].ZipCode = v.ZipCode.String
		addressPb[i].IsDefault = v.IsDefault
	}

	return addressPb
}
