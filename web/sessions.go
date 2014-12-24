package web

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	)

var cookie = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32)
	)

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func SetSession(response http.ResponseWriter) {

}

func ClearSession(response http.ResponseWriter) {

}

func LoginHandler(response http.ResponseWriter, request *http.Request) {

}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {

}