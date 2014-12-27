package web

import(
	"fmt"
	"net/http"
	_ "encoding/json"
	// "github.com/lib/pq"
	// "database/sql"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/sessions"
	_ "github.com/gorilla/securecookie"
	_ "readywater/token/token"
)


///// User /////
//Read
func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	user := GetUser()
	userJson,_ := toJson(user)
	fmt.Println("Called user",userJson)
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
	user := GetUser()
	tokens := user.GetAllTokens()
	tkJson,_ := toJson(tokens)
	fmt.Println("CallingUserToken",tkJson)
	fmt.Fprintf(res, tkJson)
}
// Get a token
func GetTokenHandler(res http.ResponseWriter, req *http.Request) {
	user := GetUser()
	tokens := user.GetAllTokens()
	tkJson,_ := toJson(tokens)
	fmt.Println("CallingTokenDetails",tkJson)
	fmt.Fprintf(res, tkJson)
}

func GetMostRecentEntriesHandler(res http.ResponseWriter, req *http.Request) {
	// user := GetUser()
	// tokens := user.tokens
	// GetMostRecentEntries()
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

