package token

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateBiscuitToken(t *testing.T) {
	email := gofakeit.Email()
	role := gofakeit.Word()
	keyPair, err := CreateKey(gofakeit.Word())

	token, payload, err := keyPair.CreateBiscuitToken(email, role)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	require.Equal(t, email, payload.Email)
	require.Equal(t, role, payload.Role)
	require.NotEmpty(t, payload.ID)
	require.NotEmpty(t, payload.IssuedAt)
	require.NotEmpty(t, payload.ExpiredAt)

	newPayload, err := keyPair.VerifyBiscuitToken(token, payload)
	require.NoError(t, err)
	require.NotEmpty(t, newPayload)
}
