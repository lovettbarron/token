package web

import(
	"fmt"
	"encoding/json"
	"readywater/token/token"
	"database/sql"
	 _ "github.com/lib/pq"
	"log"

)

const (
	
	dbType = "postgres"
	dbName = "token_test"
	dbUser = "andrewlb"
	dbPass = "password"
	dbUrl = "localhost"

	salt = "Yp2iD6PcTwB6upati0bPw314GrFWhUy90BIvbJTj5ETbbE8CoViDDGsJS6YHMOBq4VlwW3V00GWUMbbV"
	)



// Fake user until DB setup
var test = false;
var u = token.NewUser(0,0,0,"test0","test0","test0@test","Testing0")

func GetUser() *token.User {
	if !test {		
		u.NewToken("Token1",4,"a day")
		u.NewToken("Token1",10,"a week")
		u.NewToken("Token1",1,"a month")
		test=true
	}

	fmt.Println("GetUser() called:",u)
	return u
}

func toJson(data interface{}) (string, error) {
	// fmt.Println("toJson got",data)
	dataJson,err := json.Marshal(data)
	// fmt.Println("toJson made",string(dataJson))
	if err != nil {
		fmt.Println(err)
		return "error",err
	}
	return string(dataJson),nil
}

func ConnectToDB() (db *sql.DB)	 {
	db, err := sql.Open("postgres","host=localhost dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DisconnectFromDB(db *sql.DB) int {
	return -1
}

func Q(string, db *sql.DB) {

}

/*
func hashPassword(username, password string) string {

	ps := []string{password, username, salt}

	// hashed_password
	hash := sha256.New()
	hash.Write([]byte(strings.Join(ps, "-")))
	md := hash.Sum(nil)
	hashed_password := hex.EncodeToString(md)

	return hashed_password
}

// from: https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64

// See recommended IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func randString(n int) string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b % byte(len(alphanum))]
    }
    return string(bytes)
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext
}

func decrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return decodeBase64(string(text))
}
*/