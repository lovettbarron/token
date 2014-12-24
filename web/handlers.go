package web

import(
	"fmt"
	"net/http"
	"encoding/json"
	// "github.com/lib/pq"
	// "database/sql"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/sessions"
	_ "github.com/gorilla/securecookie"
	"readywater/token/token"
)

// Fake user until DB setup
var u = token.NewUser(0,0,0,"test0","test0","test0@test","Testing0")



func GetUser() *token.User {
	return u
}

func toJson(data interface{}) (string, error) {
	dataJson,err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "error",err
	}
	return string(dataJson),nil
}

/////// General //////


///// User /////
//Read
func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	user := GetUser()
	userJson,_ := toJson(user)
	fmt.Fprintf(res, userJson)
}

//update/create
func UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
	user := GetUser()
	userJson,_ := toJson(user)
	fmt.Fprintf(res, "Updated", userJson) // Remove
}

//Delete
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	user := GetUser()
	userJson,_ := toJson(user)
	fmt.Fprintf(res, "Deleted", userJson)
}

///// Token //////
//New Token
func CreateTokenHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()
	// user.NewToken(req.)
	fmt.Fprintf(res, "CreatedToken")
}

//Get All
func GetTokenListHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()

}
// Get a token
func GetTokenHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()
}

// Remove Token
func DeleteTokenHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()
}


/// Entries ///
// Use Token
func UseTokenHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()
}

// GetMostRecent

