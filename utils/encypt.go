package utils

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func IsEqual(encryptStr string, toEncryptStr string) bool {
	encryptedStr := fmt.Sprintf("%x", md5.Sum([]byte(toEncryptStr)))
	return encryptStr == encryptedStr
}

func Md5(toEncryptStr string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(toEncryptStr)))
}

func EncryptPassword(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ValidatePassword(origin, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(origin))
}
