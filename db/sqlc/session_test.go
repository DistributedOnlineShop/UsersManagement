package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"UsersManagement/util"
)

func TestCreateToken(t *testing.T) {
	user := CreateRandomUser(t)

	data := CreateTokenParams{
		TokenID:   util.CreateUUID(),
		Email:     user.Email,
		Token:     util.CreateUUID(),
		Status:    util.GenerateSessionStatus(),
		ExpiresAt: util.GenerateDate(),
	}

	token, err := testStore.CreateToken(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.Equal(t, data.TokenID, token.TokenID)
	require.Equal(t, data.Email, token.Email)
	require.Equal(t, data.Token, token.Token)
	require.Equal(t, data.Status, token.Status)
	require.NotEmpty(t, token.CreatedAt)
}
