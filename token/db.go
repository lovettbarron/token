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
	var user *User
	tdb.QueryRow("SELECT DISTINCT * FROM user WHERE id = $1",id).Scan(&user)
	return user,nil
}