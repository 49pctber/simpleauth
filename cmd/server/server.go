package main

import (
	"fmt"
	"net/http"

	"github.com/49pctber/jwtauth"
)

func main() {

	config := jwtauth.AuthConfig{}
	err := config.ReadFromFile()
	if err != nil {
		new_config, err := jwtauth.NewAuthConfig()
		if err != nil {
			panic(err)
		}
		config = *new_config

		user, err := jwtauth.NewUser("bryan", "yo mama")
		if err != nil {
			panic(err)
		}

		config.Users = append(config.Users, *user)
		config.WriteToFile()
	}

	server := http.Server{
		Addr: ":8080",
	}

	// TODO implement login page

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err = r.Cookie("auth_token")

		if err != nil {
			token_string, err := jwtauth.GenerateJWT("user", config.Secret)
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

			fmt.Fprintf(w, "Check your cookies. You just got a new one.")
		} else {
			fmt.Fprintf(w, "You already have a cookie!")
		}

	})

	fmt.Printf("Serving at %s\n", server.Addr)
	fmt.Println(server.ListenAndServe())
}
