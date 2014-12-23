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
        newUser := token.NewUser(0,0,0,"test","test","test@test","Testing")
        // userid, pwHash, pwSalt int32, username, fullname, email, intention string
        fmt.Println(newUser)

        newToken := token.NewToken("Build Token",3)
        fmt.Println(newToken)

        _ = newUser.AppendToken(newToken)
		fmt.Println(newUser)        
}



