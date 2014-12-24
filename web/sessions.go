package web

import(
	"net/http"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	)

// Snagged most of this from: https://gist.github.com/mschoebel/9398202

var cookie = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

func getUserName(req *http.Request) (userName string) {
	userName = "test0"
	// if cookie, err := req.Cookie("session"); err == nil {
	// 	cookieValue := make(map[string]string)
	// 	if err = cookie.Decode("session", cookie.Value, &cookieValue); err == nil {
	// 		userName = cookieValue["name"]
	// 	}
	// }
	return userName
}

func SetSession(userName string, res http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookie.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(res, cookie)
	}
}

func ClearSession(res http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {

}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {

}

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	}