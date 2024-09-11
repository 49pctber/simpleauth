package simpleauth

import (
	"context"
	"html/template"
	"log"
	"net/http"

	_ "embed"
)

type contextKey string

const usernameKey contextKey = "username"
const redirectKey contextKey = "redirect"

var authHandler http.HandlerFunc = http.HandlerFunc(logonHandleFunc) // for handling authentication

//go:embed login_form.tmpl
var login_form string

func GetUser(r *http.Request) string {
	uname := r.Context().Value(usernameKey)
	if uname == nil {
		return ""
	}

	return uname.(string)
}

func logonHandleFunc(w http.ResponseWriter, r *http.Request) {

	redirect := "/"
	message := "Please log in to continue"

	rd := r.Context().Value(redirectKey)
	if rd != nil {
		redirect = rd.(string)
	}

	if r.Method == http.MethodPost {

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		redirect = r.PostFormValue("redirect")

		u := FindUser(username)

		if u != nil {
			if u.ValidatePassword(password) {

				token_string, err := GenerateJWT(username, config.Secret)
				if err != nil {
					log.Println(err)
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

				http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)

			}
		}

		message = "Incorrect username or password"
	}

	tmpl, err := template.New("login_form").Parse(login_form)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct{ Message, Redirect string }{Message: message, Redirect: redirect})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		ctx = context.WithValue(ctx, redirectKey, r.URL.Path)
		ctx = context.WithValue(ctx, usernameKey, nil)

		if !config.IsInitialized() {
			log.Println(ErrsimpleauthNotConfigured)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		cookie, err := r.Cookie("auth_token")

		if err != nil {
			// user doesn't have authentication cookie
			authHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		token, err := ValidateJWT(cookie.Value, config.Secret)
		if err != nil {
			// invalid token
			authHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		username, err := token.Claims.GetSubject()
		if err != nil {
			// error getting username
			authHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		u := FindUser(username)

		if u == nil {
			// username not found
			authHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx = context.WithValue(r.Context(), usernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
