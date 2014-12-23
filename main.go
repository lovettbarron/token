package main

import (
	"fmt"
	// "github.com/go-martini/martini"
	// "github.com/martini-contrib/auth"
	// "github.com/martini-contrib/sessions"
	// "os"
	// "log"
	// "net"
	// "net/http"
	// "os/signal"
	// "time"
	"readywater/token/token"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func main() {
        user0 := token.NewUser(0,0,0,"test0","test0","test0@test","Testing0")
        user1 := token.NewUser(1,0,0,"test1","test1","test1@test","Testing1")
        // userid, pwHash, pwSalt int32, username, fullname, email, intention string
        fmt.Println(user1)

        newToken := user1.NewToken("Build Token",token.GetCycle(3,"a day"))

		fmt.Println(user1)        
		fmt.println(newToken)
		fmt.Println(user1.GetLastToken())

		tokenId := user1.GetLastToken().id;
		user1.UseToken()
}



