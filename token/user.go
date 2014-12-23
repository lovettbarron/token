package token

import(
	"sync"
	"time"

)


type User struct {
	userid int64
	pwHash int64
	pwSalt int64
	created int64
	updated int64
	lastActive int64
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

// (u *User)
func NewUser(userid, pwHash, pwSalt int32, username, fullname, email, intention string) *User {

	newUser := &User{
		0,
		0,
		0,
		time.Now().Unix(),
		time.Now().Unix(),
		time.Now().Unix(),
		username,
		fullname,
		email,
		intention,
		0,
		false,
		make([]*Token,0),
		new(sync.Mutex),
	}

	// _ = newUser // Just to get to compile

	return newUser
}

func (u *User) Authenticate() int {
	return 0
}

func (u *User) UpdateInfo(name, email, intention, password string) int {
	return 0
}

func (u *User) DeleteUser() int {
	return 0
}

///////// Tokens //////////
// Create Token
func (u *User) AppendToken(token *Token) int {
	u.tokens = append(u.tokens,token)
	return 0
}

// Get List of All Tokens
func (u *User) GetAllTokens() []*Token {
	return u.tokens
}

// Get most recently made Token (top)
func (u *User) GetLastToken() *Token {
	return u.tokens[len(u.tokens)-1]
}

// Delete all Tokens
func (u *User) RemoveAllTokens() int {
	u.tokens = make([]*Token,0)
	return 0
}

// Delete token
func (u *User) RemoveToken(tokenid int32) int {
	return 0
}