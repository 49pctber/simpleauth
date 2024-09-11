package simpleauth

import "testing"

func TestPassword(t *testing.T) {

	username := "usrnme"
	password := "secure_passw0rd!"

	user, err := NewUser(username, password, false)
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

func TestUsername(t *testing.T) {
	if have, want := ValidateUsername("a"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("ab"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("abc"), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("12345678901234567890123456789012"), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("123456789012345678901234567890123"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("abc!"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("abc😂"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername(" abc"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("abc "), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("ab-c"), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ValidateUsername("AaAa00"), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
