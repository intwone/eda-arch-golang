package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	GenerateHashError = errors.New("error to generate hash")
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) Hash(value string) (*string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return nil, GenerateHashError
	}
	hashedString := string(hashed)
	return &hashedString, nil
}

func (b *BcryptHasher) Compare(value string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
