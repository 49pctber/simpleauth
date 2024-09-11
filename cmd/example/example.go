package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "embed"

	"github.com/49pctber/jwtauth"
)

//go:embed greeting.tmpl
var greeting string

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {

	username := jwtauth.GetUser(r)

	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl, err := template.New("greeting").Parse(greeting)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func main() {

	// have to specify configuration file
	err := jwtauth.Configure("config.json")
	if err != nil {
		panic(err)
	}

	// configure server to serve on port :8080
	mux := http.NewServeMux()
	mux.Handle("/", jwtauth.RequireAuthentication(http.HandlerFunc(homeHandleFunc)))

	fmt.Println("Serving on :8080")
	fmt.Println(http.ListenAndServe(":8080", mux))
}
