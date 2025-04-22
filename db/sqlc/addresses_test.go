package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"UsersManagement/util"
)

func CreateRandomAddress(t *testing.T, userId uuid.UUID) Address {
	data := CreateAddressParams{
		AddressID: util.CreateUUID(),
		UserID:    userId,
		FlatFloor: util.GenerateText(),
		Building:  util.GenerateText(),
		Street:    gofakeit.Street(),
		District:  gofakeit.City(),
		Region:    gofakeit.State(),
		ZipCode:   util.GenerateText(),
		Country:   gofakeit.Country(),
		IsDefault: gofakeit.Bool(),
	}

	address, err := testStore.CreateAddress(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, address.AddressID)
	require.Equal(t, data.UserID, address.UserID)
	require.Equal(t, data.FlatFloor, address.FlatFloor)
	require.Equal(t, data.Building, address.Building)
	require.Equal(t, data.Street, address.Street)
	require.Equal(t, data.District, address.District)
	require.Equal(t, data.Region, address.Region)
	require.Equal(t, data.ZipCode, address.ZipCode)
	require.Equal(t, data.Country, address.Country)
	require.Equal(t, data.IsDefault, address.IsDefault)
	require.NotEmpty(t, address.CreatedAt)

	return address
}

func TestCreateAddress(t *testing.T) {
	user := CreateRandomUser(t)
	CreateRandomAddress(t, user.UserID)
}

func TestDeleteAddress(t *testing.T) {
	user := CreateRandomUser(t)
	ad := CreateRandomAddress(t, user.UserID)

	data := DeleteAddressParams{
		UserID:    user.UserID,
		AddressID: ad.AddressID,
	}

	err := testStore.DeleteAddress(context.Background(), data)
	require.NoError(t, err)
}

func TestGetAddressesByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 10; i++ {
		CreateRandomAddress(t, user.UserID)
	}
	data, err := testStore.GetAddressesByUserID(context.Background(), user.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.GreaterOrEqual(t, len(data), 10)
}

func TestResetDefaultAddress(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 3; i++ {
		CreateRandomAddress(t, user.UserID)
	}

	err := testStore.SetAllAddresstoFalse(context.Background(), user.UserID)
	require.NoError(t, err)
}

func TestSetDefaultAddress(t *testing.T) {
	user := CreateRandomUser(t)
	ad := CreateRandomAddress(t, user.UserID)

	data := SetDefaultAddressParams{
		UserID:    user.UserID,
		AddressID: ad.AddressID,
	}

	err := testStore.SetDefaultAddress(context.Background(), data)
	require.NoError(t, err)
}

func TestUpdateAddress(t *testing.T) {
	user := CreateRandomUser(t)
	ad := CreateRandomAddress(t, user.UserID)

	newData := UpdateAddressParams{
		AddressID: ad.AddressID,
		UserID:    ad.UserID,
		FlatFloor: util.GenerateText(),
		Building:  util.GenerateText(),
		Street:    gofakeit.Street(),
		District:  gofakeit.City(),
		Region:    gofakeit.State(),
		ZipCode:   util.GenerateText(),
		Country:   gofakeit.Country(),
		IsDefault: gofakeit.Bool(),
	}
	err := testStore.UpdateAddress(context.Background(), newData)
	require.NoError(t, err)
}
