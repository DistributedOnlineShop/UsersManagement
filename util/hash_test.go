package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	salt := "test@example.com"
	password := "asdasdasdasda12312"

	Hash, err := Hash(password, salt)
	require.NoError(t, err)
	require.NotEmpty(t, Hash)
	require.Len(t, Hash, 32)

	verified := VerifyHashPassword(password, salt, Hash)
	require.True(t, verified)
}
