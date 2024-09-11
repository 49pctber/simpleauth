package jwtauth

import (
	"crypto/rand"
	"encoding/json"
	"os"
)

// TODO choose more appropriate location
const ConfigFilename string = "config.json"

type AuthConfig struct {
	Secret []byte
	Users  []User
}

func NewAuthConfig() (*AuthConfig, error) {

	ac := &AuthConfig{
		Secret: make([]byte, 32),
		Users:  make([]User, 0),
	}

	_, err := rand.Read(ac.Secret)
	if err != nil {
		return nil, err
	}

	return ac, nil
}

func (ac AuthConfig) WriteToFile() error {
	data, err := json.MarshalIndent(ac, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFilename, data, 0744)
}

func (ac *AuthConfig) ReadFromFile() error {

	data, err := os.ReadFile(ConfigFilename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, ac)
}
