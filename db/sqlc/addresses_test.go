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
		AddressID:  util.CreateUUID(),
		UserID:     userId,
		Address:    gofakeit.Address().Address,
		City:       gofakeit.City(),
		State:      gofakeit.State(),
		PostalCode: gofakeit.Zip(),
		Country:    gofakeit.Country(),
		IsDefault:  gofakeit.Bool(),
	}

	address, err := testStore.CreateAddress(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, address.AddressID)
	require.Equal(t, data.UserID, address.UserID)
	require.Equal(t, data.Address, address.Address)
	require.Equal(t, data.City, address.City)
	require.Equal(t, data.State, address.State)
	require.Equal(t, data.PostalCode, address.PostalCode)
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
	ad := CreateRandomAddress(t, user.UserID)

	add, err := testStore.ResetDefaultAddress(context.Background(), ad.AddressID)
	require.NoError(t, err)
	require.NotEmpty(t, add)
	require.False(t, add.IsDefault)
}

func TestSetDefaultAddress(t *testing.T) {
	user := CreateRandomUser(t)
	ad := CreateRandomAddress(t, user.UserID)

	add, err := testStore.SetDefaultAddress(context.Background(), ad.AddressID)
	require.NoError(t, err)
	require.NotEmpty(t, add)
	require.True(t, add.IsDefault)
}

func TestUpdateAddress(t *testing.T) {
	user := CreateRandomUser(t)
	ad := CreateRandomAddress(t, user.UserID)

	newData := UpdateAddressParams{
		AddressID:  ad.AddressID,
		Address:    gofakeit.Address().Address,
		City:       gofakeit.City(),
		State:      gofakeit.State(),
		PostalCode: gofakeit.Zip(),
		Country:    gofakeit.Country(),
		IsDefault:  gofakeit.Bool(),
	}
	updated, err := testStore.UpdateAddress(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, newData.AddressID, updated.AddressID)
	require.Equal(t, newData.Address, updated.Address)
	require.Equal(t, newData.City, updated.City)
	require.Equal(t, newData.State, updated.State)
	require.Equal(t, newData.PostalCode, updated.PostalCode)
	require.Equal(t, newData.Country, updated.Country)
	require.Equal(t, newData.IsDefault, updated.IsDefault)
	require.NotEmpty(t, updated.UpdatedAt)
}
