package services_test

import (
	"testing"

	cryptography "github.com/intwone/eda-arch-golang/internal/private_cryptography/services"
	"github.com/stretchr/testify/require"
)

func TestJWTCryptography_Encrypt(t *testing.T) {
	t.Run("should encrypt a value", func(t *testing.T) {
		secret := "secret_test"
		value := "test123"
		result, err := cryptography.NewJWTCryptography(secret).Encrypt(value)

		require.Nil(t, err)
		require.Equal(t, true, len(*result) > 0)
	})

	t.Run("should decrypt a value", func(t *testing.T) {
		secret := "secret_test"
		value := "test123"
		crypto := cryptography.NewJWTCryptography(secret)
		result1, _ := crypto.Encrypt(value)
		result2, err := crypto.Decrypt(*result1)

		require.Nil(t, err)
		require.Equal(t, value, *result2)
	})

}
