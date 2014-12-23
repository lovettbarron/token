package token

import(
	"sync"
	"time"

)


type User struct {
	userid int32
	pwHash int32
	pwSalt int32
	created int32
	updated int32
	lastActive int32
	userName string
	fullName string
	email string
	intention string
	score int
	disable bool
	tokens []*Token
	mutex *sync.Mutex
}

// Is this chunk necessary?
// func newUser() *User {
// 	return &User{
// 		make([]*Token,0)
// 		new(sync.Mutex)
// 	}
// }

func (u *User) defineNewUser(userid, pwHash, pwSalt int32, username, fullname, email, intention string) int {

	newUser := &User{
		0
		0,
		0,
		time.Now(),
		time.Now()
		time.Now(),
		username
		fullname
		email
		intention
		0
		false
		make([]*Token,0)
	}

	return 0
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