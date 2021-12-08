package utils

import (
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
