package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"readywater/token/token"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func main() {

}

func test() {
        user0 := token.NewUser(0,0,0,"test0","test0","test0@test","Testing0")
        user1 := token.NewUser(1,0,0,"test1","test1","test1@test","Testing1")
        // userid, pwHash, pwSalt int32, username, fullname, email, intention string
        fmt.Println(user0)
        fmt.Println(user1)

        newToken := user1.NewToken("Build Token",token.GetCycle(3,"a day"))

		fmt.Println(user1)        
		fmt.Println(newToken)
		fmt.Println(user1.GetLastToken())

		tokenId := user1.GetLastToken()
		tokenId.UseToken()
		tokenId.UseToken()
		tokenId.UseToken()
		tokenId.UseToken()
		tokenId.UseToken()
}

