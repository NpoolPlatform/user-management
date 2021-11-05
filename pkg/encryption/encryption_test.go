package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	// 1. generate a salt.
	salt := Salt()
	fmt.Println("salt is", salt)

	// 2. input a password.
	truePassword := "lpz990627"

	// 3. encrypte password.
	enPass, err := EncryptePassword(truePassword, salt)
	fmt.Println("en pass is", enPass)
	assert.Nil(t, err)

	// 4. mock user input pasword
	inputPass := "lpz990627"
	err = VerifyUserPassword(inputPass, enPass, salt)
	assert.Nil(t, err)
}
