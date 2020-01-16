package helpers

import (
	"time"

	"github.com/elithrar/simple-scrypt"
)

//GenerateFromPassword 返回加密后的密码
func GenerateFromPassword(password []byte) ([]byte, error) {
	params, err := scrypt.Calibrate(500*time.Millisecond, 64, scrypt.Params{})
	if err != nil {
		return nil, err
	}

	return scrypt.GenerateFromPassword(password, params)
}

//CompareHashAndPassword 比较加密后密码和原密码
func CompareHashAndPassword(hash []byte, password []byte) error {
	return scrypt.CompareHashAndPassword(hash, password)
}
