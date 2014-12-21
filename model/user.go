package token

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
}


func newUser() *User {

}

func (u *User) Authenticate() int {

}

func (u *User) UpdateInfo(name, email, intention, password string) int {

}

func (u *User) DeleteUser() int {

}

