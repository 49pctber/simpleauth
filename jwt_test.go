package simpleauth

import (
	"testing"
)

func TestValidation(t *testing.T) {

	secret_key := []byte("secret_bytes")
	username := "bryan"

	tokenstring, err := GenerateJWT(username, secret_key)
	if err != nil {
		t.Fatal(err)
	}

	token, err := ValidateJWT(tokenstring, secret_key)
	if err != nil {
		t.Fatal(err)
	}

	if sub, err := token.Claims.GetSubject(); err != nil || sub != username {
		t.Fatal("usernames don't match")
	}

}
