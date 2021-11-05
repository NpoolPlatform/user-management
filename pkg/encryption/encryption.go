package encryption

import (
	"crypto/md5"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

// Lower 随机生成size个小写字母
func Upper(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{65, 90}
	result := make([]byte, 26)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(26))
	}
	return result
}

// Number 随机生成size个数字
func Number(size int) []byte {
	if size <= 0 || size > 10 {
		size = 10
	}
	warehouse := []int{48, 57}
	result := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(9))
	}
	return result
}

// Lower 随机生成size个小写字母
func Lower(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{97, 122}
	result := make([]byte, 26)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + rand.Intn(26))
	}
	return result
}

// Salt 生成一个盐值
func Salt() string {
	// 按需要生成字符串
	var result string
	lowers := string(Lower(6))
	result += lowers
	numbers := string(Number(6))
	result += numbers
	uppers := string(Upper(6))
	result += uppers

	return result
}

func generateSaltPassword(password, salt string) []byte {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))

	return m5.Sum(nil)
}

func EncryptePassword(password, salt string) (string, error) {
	saltPassword := generateSaltPassword(password, salt)

	hash, err := bcrypt.GenerateFromPassword(saltPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", xerrors.Errorf("generate hash password error: %v", err)
	}

	return string(hash), nil
}

func VerifyUserPassword(inputPassword, password, salt string) error {
	saltPassword := generateSaltPassword(inputPassword, salt)

	err := bcrypt.CompareHashAndPassword([]byte(password), saltPassword)
	if err != nil {
		return xerrors.Errorf("input password is wrong!")
	}
	return nil
}
