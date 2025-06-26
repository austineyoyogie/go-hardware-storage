package utils

import (
	"math/rand"
	"time"
	"strings"
	"golang.org/x/crypto/bcrypt"

)
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(strlen int) string {
	const chars = "2a$12$cOfvxNj0xiZYs063N2Kygu2K49mPSnzH4K2vjgbhZMTxuGldov57e"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func IsTitle(value string) string  {
	return strings.Title(value)
}
func IsToLower(value string) string {
	return strings.ToLower(value)
}

func BcryptHash(password string) ([]byte, error) {
	const hashCost int = 12
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func CompareHashPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
