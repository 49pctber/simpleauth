package jwtauth

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"os"
)

var config AuthConfig

var ErrJwtauthNotConfigured error = errors.New("jwtauth has not been configured yet")

type AuthConfig struct {
	initialized bool
	Filename    string
	Secret      []byte
	Users       []User
}

func Configure(filename string) error {
	config.Filename = filename
	err := config.ReadFromFile()
	if err != nil {
		return err
	}
	config.initialized = true
	return nil
}

func NewAuthConfig(filename string) error {

	config = AuthConfig{
		initialized: true,
		Filename:    filename,
		Secret:      make([]byte, 32),
		Users:       make([]User, 0),
	}

	_, err := rand.Read(config.Secret)
	if err != nil {
		return err
	}

	return config.WriteToFile()
}

func AddUser(username, password string) error {
	return config.AddUser(username, password)
}

func (ac AuthConfig) IsInitialized() bool {
	return ac.initialized
}

func (ac AuthConfig) WriteToFile() error {
	if !ac.IsInitialized() {
		return ErrJwtauthNotConfigured
	}

	data, err := json.MarshalIndent(ac, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(config.Filename, data, 0744)
}

func (ac *AuthConfig) ReadFromFile() error {

	data, err := os.ReadFile(config.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, ac)
	if err != nil {
		return err
	}

	ac.initialized = true
	return nil
}

func (ac *AuthConfig) AddUser(username, password string) error {
	if !ac.IsInitialized() {
		return ErrJwtauthNotConfigured
	}

	user, err := NewUser(username, password)
	if err != nil {
		panic(err)
	}

	for i, u := range config.Users {
		if u.Username == username {
			config.Users[i].PasswordHash = user.PasswordHash
			config.Users[i].Salt = user.Salt
			return config.WriteToFile()
		}
	}

	config.Users = append(config.Users, *user)
	return config.WriteToFile()
}
