package token

import(
	"sync"
	"time"

)


type User struct {
	id int64			`json:"id"`
	pwHash int64		`json:"-"`
	pwSalt int64		`json:"-"`
	created int64		`json:"created"`
	updated int64		`json:"updated"`
	lastActive int64	`json:"lastActive"`
	userName string		`json:"userName"`
	fullName string		`json:"fullName"`
	email string		`json:"email"`
	intention string	`json:"intention"`
	score int 			`json:"score"`
	disable bool		`json:"-"`
	tokens []*Token     `json:"-"`
	mutex *sync.Mutex 	`json:"-"`
}

func NewUser(userid, pwHash, pwSalt int64, username, fullname, email, intention string) *User {

	newUser := &User{
		userid,
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
// Generate Token for this user
func (u *User) NewToken(title string, quant int64, interval string) *Token {
	token := &Token{
		int64(len(u.tokens)),
		u.id,
		title,
		quant,
		interval,
		time.Now().Unix(),
		make([]*TokenEntry, 0),
		new(sync.Mutex),
	}
	u.tokens = append(u.tokens,token)
	return token
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