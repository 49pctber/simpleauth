package jwtauth

import (
	"bytes"
	"crypto"
	"crypto/rand"
)

type User struct {
	Id           []byte
	Username     string
	PasswordHash []byte
	Salt         []byte
	Permissions  []string
}

func NewUser(username, password string) (*User, error) {
	user := &User{
		Username:     username,
		Id:           make([]byte, 12),
		Salt:         make([]byte, 32),
		PasswordHash: make([]byte, 0),
		Permissions:  make([]string, 0),
	}

	_, err := rand.Read(user.Id)
	if err != nil {
		return nil, err
	}

	_, err = rand.Read(user.Salt)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = user.HashPassword(password)

	return user, nil
}

func (user User) HashPassword(password string) []byte {
	h := crypto.SHA256.New()
	h.Write(user.Salt)
	h.Write([]byte(password))
	return h.Sum(nil)
}

func (user User) ValidatePassword(password string) bool {
	result := user.HashPassword(password)
	return bytes.Equal(result, user.PasswordHash)
}
