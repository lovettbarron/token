package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	// "github.com/gorilla/sessions"
	// "github.com/gorilla/securecookie"
	// "readywater/token/token"
	"readywater/token/web"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Port string = ":8000"
)

var r = mux.NewRouter()
var t = r.PathPrefix("/token").Subrouter()
var u = r.PathPrefix("/user").Subrouter()

func main() {
	r.HandleFunc("/",web.IndexHandler)
	r.HandleFunc("/login", web.LoginHandler)
	r.HandleFunc("/logout", web.LogoutHandler)

	// User Handlers
	u.HandleFunc("/",web.GetUserHandler).Methods("GET")
	u.HandleFunc("/{data}",web.UpdateUserHandler).Methods("POST")
	u.HandleFunc("/",web.DeleteUserHandler).Methods("DELETE")

	// Token Handlers
	t.HandleFunc("/", web.GetTokenListHandler).Methods("GET")
	t.HandleFunc("/", web.CreateTokenHandler).Methods("POST").
		Queries(
			"title","{title:[A-Za-z]|[0-9]|_+}",
			"quant","{quant:[0-9]+}",
			"interval","{interval:[A-Za-z]|[0-9]|_+}",
			)
	// Individual
	t.HandleFunc("/{id}", web.GetTokenHandler).Methods("GET")
	t.HandleFunc("/{id}", web.UseTokenHandler).Methods("POST")
	t.HandleFunc("/{id}", web.DeleteTokenHandler).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server running at port",Port)
	http.ListenAndServe(Port, nil)
}