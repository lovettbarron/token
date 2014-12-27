package db

import (
	"fmt"
	"database/sql"
	 _ "github.com/lib/pq"
	"log"
)

/*
type TokenEntry struct {
	Id int64		`json:"id" db:"id"`
	Tokenid int64	`json:"tokenid" db:"tokenid
	Timestamp int64	`json:"timestamp" db:"times
	Value float32	`json:"value" db:"value"`
}

type Token struct {
	Id int64				`json:"id" db:"id"`
	Userid int64			`json:"-" db:"useri
	Title string			`json:"title" db:"t
	Quantity int64 			`json:"quantity" db
	Interval string 		`json:"interval" db
	Start int64 			`json:"start" db:"s
	tokens []*TokenEntry	`json:"-" db:"-"`
	mutex *sync.Mutex		`json:"-" db:"-"`
}

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
}
*/



func main() {
	/*
CREATE USER token WITH PASSWORD 'token_test';
CREATE DATABASE token OWNER token;

CREATE TABLE "user" ( id bigserial primary key,
	pwhash int NOT NULL,
	pwsalt int NOT NULL,
	created timestamp NOT NULL,	
	updated timestamp NOT NULL,
	lastactive timestamp NOT NULL,
	username text NOT NULL,
	fullname text,
	email varchar(20) NOT NULL,
	intention text,
	score int,
	disable bool
	);

CREATE TABLE "token" ( id bigserial primary key,
	userid bigint NOT NULL,
	title text NOT NULL,
	quantity int NOT NULL,
	interval varchar(20) NOT NULL,
	start timestamp NOT NULL
	);

CREATE TABLE "tokenentry" (
	id bigserial primary key,
	tokenid bigint NOT NULL,
	timestamp timestamp NOT NULL,
	value real NOT NULL
	);

}




	*/
}