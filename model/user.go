package token

import(
	"token"
)

type User struct {
	userid int32
	pwhash int32
	pwsalt int32
	created int32
	updated int32
	lastactive int32
	username string
	fullname string
	email string
	intention string
	score int
	disable bool
	tokens []*Tokens
	mutex *sync.Mutex
}


func newUser() *User {

}

func (u *User) Authenticate() int {

}

func (u *User) UpdateInfo(name, email, intention, password string) int {

}

func (u *User) DeleteUser() int {

}

///////// Tokens //////////
// Get List of All Tokens
func (t *Token) GetAllTokens() []*Token {

}

// Get most recently made Token (top)
func (t *Token) GetLastToken() *TokenEntry {

}

// Delete all Tokens
func (t *Token) RemoveAllTokens() int {

}

// Delete token
func (t *Token) RemoveToken(tokenid int32) int {

}