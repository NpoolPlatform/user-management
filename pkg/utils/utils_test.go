package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	username := "123  dsada"
	match := RegexpUsername(username)
	assert.Equal(t, false, match)

	username = "12312312312312321312"
	match = RegexpUsername(username)
	assert.Equal(t, false, match)

	username = "242424242"
	match = RegexpUsername(username)
	assert.Equal(t, false, match)

	username = "jhdasio@jdsa.com"
	match = RegexpUsername(username)
	assert.Equal(t, false, match)

	username = "jdisa8912371"
	match = RegexpUsername(username)
	assert.Equal(t, true, match)

	username, err := GenerateUsername()
	assert.Nil(t, err)
	match = RegexpUsername(username)
	assert.Equal(t, true, match)

	passwprd := "crasd12313"
	match = RegexpPassword(passwprd)
	assert.Equal(t, true, match)

	passwprd = "12312312312"
	match = RegexpPassword(passwprd)
	assert.Equal(t, false, match)

	passwprd = "dasdsadasdasdas"
	match = RegexpPassword(passwprd)
	assert.Equal(t, false, match)

	passwprd = "ojidsoaWJ894342"
	match = RegexpPassword(passwprd)
	assert.Equal(t, true, match)
}
