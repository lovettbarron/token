package token

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func main() {
        martiniClassic := martini.Classic()
        token := token.NewToken();
        webservice.RegisterWebService(token, martiniClassic)
        martiniClassic.Run()
}