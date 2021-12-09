package utils

import (
	"net/mail"
	"regexp"

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
	"golang.org/x/xerrors"
)

func GenerateUsername() (string, error) {
	vc := vcgen.New(&vcgen.Generator{
		Count:   1,
		Pattern: "##########",
		Charset: "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM",
	})
	codes, err := vc.Run()
	if err != nil {
		return "", xerrors.Errorf("fail to run invitation code generator: %v", err)
	}
	return (*codes)[0], nil
}

func RegexpUsername(username string) bool {
	if b, err := regexp.MatchString("^[0-9]*$", username); b {
		if err == nil {
			return false
		}
		return false
	}

	_, err := mail.ParseAddress(username)
	return err != nil
}
