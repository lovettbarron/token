package token

import(
	"sync"
	"time"
	"fmt"
	"math/rand"
)

type User struct {
	Id int64			`json:"id" db:"id"`
	PwHash int64		`json:"-" db:"pwhash"`
	PwSalt int64		`json:"-" db:"pwsalt"`
	Created int64		`json:"created" db:"created"`
	Updated int64		`json:"updated" db:"updated"`
	LastActive int64	`json:"lastActive" db:"lastactive"`
	UserName string		`json:"userName" db:"username"`
	FullName string		`json:"fullName" db:"fullname"`
	Email string		`json:"email" db:"email"`
	Intention string	`json:"intention" db:"intention"`
	Score int64			`json:"score" db:"score"`
	Disable bool		`json:"-" db:"disable"`
	Tokens []*Token     `json:"-" db:"-"`
	Mutex *sync.Mutex 	`json:"-" db:"-"`
}

func CreateEmptyUser() *User {

	newUser := &User{
		0,
		0,
		0,
		time.Now().Unix(),
		time.Now().Unix(),
		time.Now().Unix(),
		"",
		"",
		"",
		"",
		0,
		false,
		make([]*Token,0),
		new(sync.Mutex),
	}

	return newUser
}

func NewUser(userid, pwHash, pwSalt int64, username, fullname, email, intention string) int64 {

	fmt.Println("Making new user")

	theId := rand.Int63()

	newUser := &User{
		theId,
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

	if tdb == nil  { return -1 }

	err := tdb.QueryRow("insert into public.user (id, pwhash, pwsalt,created,updated,lastactive,username,fullname,email,intention,score,disable) values ( $1,$2,$3,to_timestamp($4),to_timestamp($5),to_timestamp($6),$7,$8,$9,$10,$11,$12) RETURNING id",theId,newUser.PwHash, newUser.PwSalt, newUser.Created,newUser.Updated, newUser.LastActive, newUser.UserName,newUser.FullName, newUser.Email, newUser.Intention,newUser.Score, newUser.Disable ).Scan(&theId)

	if err != nil {
	    fmt.Println(err)
	    return -1
	}

	return theId
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
func (u *User) NewToken(title string, quant int64, interval string) int64 {
	// token := &Token{
	// 	int64(len(u.Tokens)),
	// 	u.Id,
	// 	title,
	// 	quant,
	// 	interval,
	// 	time.Now().Unix(),
	// 	make([]*TokenEntry, 0),
	// 	new(sync.Mutex),
	// }

	tokenId := rand.Int63()

	err := tdb.QueryRow("insert into public.token (id, userid, title, quantity, interval, start) values ( $1,$2,$3,$4,$5,to_timestamp($6)) RETURNING id",tokenId,u.Id, title, quant, interval, time.Now().Unix()).Scan(&tokenId)

	if err != nil {
	    fmt.Println(err)
	    return -1
	}

	return tokenId
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