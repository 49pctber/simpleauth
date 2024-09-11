package jwtauth

import "testing"

func TestPassword(t *testing.T) {

	username := "usrnme"
	password := "secure_passw0rd!"

	user, err := NewUser(username, password)
	if err != nil {
		t.Fatal(err)
	}

	if !user.ValidatePassword(password) {
		t.Error("password validation failed with correct password")
	}

	if user.ValidatePassword("not my password") {
		t.Error("password validation succeeded with incorrect password")
	}

	user.Salt[0] ^= 0x80 // use incorrect salt

	if user.ValidatePassword(password) {
		t.Error("password validation succeeded with incorrect salt")
	}

	user.Salt[0] ^= 0x0f // use incorrect salt

	if user.ValidatePassword(password) {
		t.Error("password validation succeeded with incorrect salt")
	}

	user.Salt[0] ^= 0x8f // use correct salt again

	if !user.ValidatePassword(password) {
		t.Error("password validation failed")
	}
}
