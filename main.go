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

func test() {
  //       user0 := token.NewUser(0,0,0,"test0","test0","test0@test","Testing0")
  //       user1 := token.NewUser(1,0,0,"test1","test1","test1@test","Testing1")
  //       // userid, pwHash, pwSalt int32, username, fullname, email, intention string
  //       fmt.Println(user0)
  //       fmt.Println(user1)

  //       newToken := user1.NewToken("Build Token",token.GetCycle(3,"a day"))

		// fmt.Println(user1)        
		// fmt.Println(newToken)
		// fmt.Println(user1.GetLastToken())

		// tokenId := user1.GetLastToken()
		// tokenId.UseToken()
		// tokenId.UseToken()
		// tokenId.UseToken()
		// tokenId.UseToken()
		// tokenId.UseToken()
}

