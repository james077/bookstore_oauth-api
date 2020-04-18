package access_token

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "El tiempo de expiración deberi se de 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "El token de acceso nuevo no debe expirar")
	assert.EqualValues(t, "", at.AccessToken, "El nuevo token de acceso no debería tener una identificación de token de acceso definido un id de token de acceso")
	assert.True(t, at.UserId == 0, "El nuevo token de acceso no debería tener un id de usuario asociado")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "El token de acceso vacío debe caducar por defecto")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "El token de acceso que expira dentro de tres horas a partir de ahora NO debe expirar")
}
