package web

import(
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"readywater/token/token"
)

///// User /////
//Read
func GetUserHandler(res http.ResponseWriter, req *http.Request) {

}

//update/create
func UpdateUserHandler(res http.ResponseWriter, req *http.Request) {

}

//Delete
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	
}

///// Token //////
//New Token
func CreateTokenHandler(res http.ResponseWriter, req *http.Request) {

}

//Get All
func GetTokenListHandler(res http.ResponseWriter, req *http.Request) {

}
// Get a token
func GetTokenHandler(res http.ResponseWriter, req *http.Request) {

}

// Remove Token
func DeleteTokenHandler(res http.ResponseWriter, req *http.Request) {

}


/// Entries ///
// Use Token
func UseTokenHandler(res http.ResponseWriter, req *http.Request) {

}

// GetMostRecent

