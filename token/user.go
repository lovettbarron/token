package token

import(
	"sync"
	"time"

)

type User struct {
	Id int64			`json:"id"`
	PwHash int64		`json:"-"`
	PwSalt int64		`json:"-"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
	LastActive int64	`json:"lastActive"`
	UserName string		`json:"userName"`
	FullName string		`json:"fullName"`
	Email string		`json:"email"`
	Intention string	`json:"intention"`
	Score int64			`json:"score"`
	Disable bool		`json:"-"`
	Tokens []*Token     `json:"-"`
	Mutex *sync.Mutex 	`json:"-"`
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
		int64(len(u.Tokens)),
		u.Id,
		title,
		quant,
		interval,
		time.Now().Unix(),
		make([]*TokenEntry, 0),
		new(sync.Mutex),
	}
	u.Tokens = append(u.Tokens,token)
	return token
}

// Get List of All Tokens
func (u *User) GetAllTokens() []*Token {
	return u.Tokens
}

// Get most recently made Token (top)
func (u *User) GetLastToken() *Token {
	return u.Tokens[len(u.Tokens)-1]
}

// Retrieve specific Token
func (u *User) GetToken(id int64) *Token {
	for _,t := range u.Tokens {
		if t.Id == id {
			return t
		}
	}
	return nil
}

// Delete all Tokens
func (u *User) RemoveAllTokens() int {
	u.Tokens = make([]*Token,0)
	return 0
}

// Delete token
func (u *User) RemoveToken(tokenid int32) int {
	return 0
}