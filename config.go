package simpleauth

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"os"
	"slices"
)

var config AuthConfig

const DefaultConfigFilename string = "simpleauth.json"

var ErrsimpleauthNotConfigured error = errors.New("simpleauth has not been configured yet")

type AuthConfig struct {
	initialized bool
	filename    string
	Secret      []byte
	Users       []User
}

func Configure(filename string) error {
	config.filename = filename
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
		filename:    filename,
		Secret:      make([]byte, 32),
		Users:       make([]User, 0),
	}

	_, err := rand.Read(config.Secret)
	if err != nil {
		return err
	}

	return config.WriteToFile()
}

func FindUser(username string) *User {
	for _, u := range config.Users {
		if u.Username == username {
			return &u
		}
	}
	return nil
}

func GetUsernames() []string {
	usernames := make([]string, len(config.Users))

	for i, u := range config.Users {
		usernames[i] = u.Username
	}

	slices.Sort(usernames)

	return usernames
}

func AddUser(username, password string, admin bool) error {
	return config.AddUser(username, password, admin)
}

func DeleteUser(username string) error {
	for i, u := range config.Users {
		if u.Username == username {
			if i == len(config.Users)-1 {
				config.Users = config.Users[:len(config.Users)-1]
			} else {
				config.Users = append(config.Users[:i], config.Users[i+1:]...)
			}
			return config.WriteToFile()
		}
	}
	return errors.New("user not found")
}

func (ac AuthConfig) IsInitialized() bool {
	return ac.initialized
}

func (ac AuthConfig) WriteToFile() error {
	if !ac.IsInitialized() {
		return ErrsimpleauthNotConfigured
	}

	data, err := json.MarshalIndent(ac, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(config.filename, data, 0744)
}

func (ac *AuthConfig) ReadFromFile() error {

	data, err := os.ReadFile(config.filename)
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

func (ac *AuthConfig) AddUser(username, password string, admin bool) error {
	if !ac.IsInitialized() {
		return ErrsimpleauthNotConfigured
	}

	user, err := NewUser(username, password, admin)
	if err != nil {
		panic(err)
	}

	u := FindUser(username)
	if u != nil {
		u.PasswordHash = user.PasswordHash
		u.Salt = user.Salt
	} else {
		config.Users = append(config.Users, *user)
	}

	return config.WriteToFile()
}
