package token

import(
	"fmt"
	"database/sql"
	 _ "github.com/lib/pq"
	"log"
	"strings"
)

const (
	
	dbType = "postgres"
	dbName = "token_test"
	dbUser = "token"
	dbPass = "tokentest"
	dbHost = "localhost"
)

var tdb *sql.DB



func ConnectToDB() {
	t:= []string{
		"host=",dbHost,
		" dbname=",dbName,
		" user=",dbUser,
		" password=",dbPass,
		" sslmode=disable",
	}
	s := strings.Join(t,"")
	var err error
	tdb, err = sql.Open(dbType,s)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	err = tdb.Ping()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println("Successfully connect to",dbName)
	fmt.Println("Successfully connect to",tdb)
	// tdb = db
}

func DisconnectFromDB() {
	err := tdb.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("Successfully closed", dbName)
}

func FetchUserFromDB(id int64) (*User,error) {
	// fmt.Println("db",tdb)
	if tdb == nil  { return nil,nil }
	var u *User
	u = CreateEmptyUser()

	rows,_ := tdb.Query("SELECT DISTINCT * FROM public.user WHERE id=($1)",id)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.Id, &u.PwHash, &u.PwSalt, &u.Created, &u.Updated, &u.LastActive, &u.UserName, &u.FullName, &u.Email, &u.Email, &u.Intention, &u.Score, &u.Disable)
		fmt.Println("User fetched ", u.Id)
	}

	return u,nil
}

func CheckUserExists(username string, hash, salt int64) int64 {
	if tdb == nil  { return -1 }
	var id int64
	id=-1

	rows,_ := tdb.Query("SELECT DISTINCT id FROM public.user WHERE username=($1) AND pwhash=($2) AND pwsalt=($3)",username, hash, salt)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id)
		fmt.Println("User fetched ", id)
	}
	return id
}