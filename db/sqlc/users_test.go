package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"UsersManagement/util"
)

func CreateRandomUser(t *testing.T) User {
	email := gofakeit.Email()
	password := gofakeit.Word()
	hash, err := util.Hash(email, password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	data := CreateUserParams{
		UserID:       util.CreateUUID(),
		FristName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		Email:        email,
		PhoneNumber:  gofakeit.Phone(),
		PasswordHash: hash,
		Status:       util.GenerateUserStatus(),
	}

	user, err := testStore.CreateUser(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, user.UserID)
	require.Equal(t, data.FristName, user.FristName)
	require.Equal(t, data.LastName, user.LastName)
	require.Equal(t, data.Email, user.Email)
	require.Equal(t, data.PasswordHash, user.PasswordHash)
	require.Equal(t, data.Status, user.Status)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user := CreateRandomUser(t)

	row, err := testStore.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, row)
}

func TestResetEmail(t *testing.T) {
	user := CreateRandomUser(t)

	newData := ResetEmailParams{
		PhoneNumber: user.PhoneNumber,
		Email:       gofakeit.Email(),
	}

	err := testStore.ResetEmail(context.Background(), newData)
	require.NoError(t, err)
}

func TestResetPassword(t *testing.T) {
	user := CreateRandomUser(t)

	password := gofakeit.Word()
	hash, err := util.Hash(user.Email, password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	newData := ResetPasswordParams{
		Email:        user.Email,
		PasswordHash: hash,
	}

	err = testStore.ResetPassword(context.Background(), newData)
	require.NoError(t, err)
}

func TestResetPhoneNumber(t *testing.T) {
	user := CreateRandomUser(t)

	newData := ResetPhoneNumberParams{
		Email:       user.Email,
		PhoneNumber: gofakeit.Phone(),
	}

	err := testStore.ResetPhoneNumber(context.Background(), newData)
	require.NoError(t, err)
}

func TestUserLogin(t *testing.T) {
	user := CreateRandomUser(t)

	hash, err := testStore.UserLogin(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
}
