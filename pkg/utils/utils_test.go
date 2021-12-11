package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	username := "12312312312312321312"
	match := RegexpUsername(username)
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
	fmt.Println("user name is", username)
	match = RegexpUsername(username)
	assert.Equal(t, true, match)
}
