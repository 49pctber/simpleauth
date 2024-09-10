package main

import (
	"fmt"
	"net/http"

	"github.com/49pctber/jwtauth"
)

func main() {

	secret_key := []byte("yo mama!!!")

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("auth_token")

		if err != nil {
			token_string, err := jwtauth.GenerateJWT("user", secret_key)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    token_string,
				HttpOnly: true, // Prevents client-side scripts from accessing the cookie
				Secure:   true, // Ensures the cookie is sent only over HTTPS
				Path:     "/",
				SameSite: http.SameSiteStrictMode, // Mitigates CSRF attacks
			})
		}

		fmt.Fprintf(w, "Check your cookies.")

	})

	fmt.Println(server.ListenAndServe())
}
