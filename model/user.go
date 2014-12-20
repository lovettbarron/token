package user

type User struct {
	userid int32
	password int32
	name string
	email string
	intention string
	score int
}


func newUser() *User {

}

func (u *User) UpdateInfo(name, email, intention, password string) int {

}

func (u *User) DeleteUser() int {

}

