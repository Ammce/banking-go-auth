package utils

import "golang.org/x/crypto/bcrypt"

func Hash(value string) string {
	byteVal := []byte(value)

	hashedValue, err := bcrypt.GenerateFromPassword(byteVal, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedValue)
}
