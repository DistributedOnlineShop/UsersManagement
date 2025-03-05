package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"UsersManagement/util"
)

func TestCreateSession(t *testing.T) {
	user := CreateRandomUser(t)

	data := CreateSessionParams{
		SessionID: util.CreateUUID(),
		Email:     user.Email,
		Token:     util.CreateUUID().String(),
		Status:    util.GenerateSessionStatus(),
		ExpiresAt: util.GenerateDate(),
	}

	token, err := testStore.CreateSession(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.Equal(t, data.SessionID, token.SessionID)
	require.Equal(t, data.Email, token.Email)
	require.Equal(t, data.Token, token.Token)
	require.Equal(t, data.Status, token.Status)
	require.NotEmpty(t, token.CreatedAt)
}
