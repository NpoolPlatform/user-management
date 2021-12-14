package utils

import (
	"net/mail"
	"regexp"
	"strings"

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
	if len(username) < 4 || len(username) > 32 {
		return false
	}

	if b, err := regexp.MatchString("^[0-9]*$", username); b {
		if err == nil {
			return false
		}
		return false
	}

	if ok := strings.Contains(username, " "); ok {
		return false
	}

	_, err := mail.ParseAddress(username)
	return err != nil
}

func RegexpPassword(password string) bool { // nolint
	if b, err := regexp.MatchString("^[0-9]*$", password); b {
		if err == nil {
			return false
		}
		return false
	}

	if b, err := regexp.MatchString("^[a-zA-Z]*$", password); b {
		if err == nil {
			return false
		}
		return false
	}

	if ok := strings.Contains(password, " "); ok {
		return false
	}

	if password == "" {
		return false
	}

	return true
}
