package dao

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	U "github.com/abdullahgumi/authentication/models"

	_ "github.com/mattn/go-sqlite3"

	"github.com/abdullahgumi/authentication/util"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "./store.db3")
	util.HandleErr(err)
	createTableQuery := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email VARCHAR(35) NOT NULL,
		password VARCHAR(32) NOT NULL,
		firstName VARCHAR(25) NOT NULL,
		lastName VARCHAR(25) NOT NULL,
		phone VARCHAR(13) NOT NULL,
		address TEXT NOT NULL,
		city VARCHAR(20) NOT NULL,
		userType VARCHAR(15) DEFAULT "client",
		lat REAL DEFAULT 0.0,
		lng REAL DEFAULT 0.0,
		status INTEGER DEFAULT 0,
		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);CREATE UNIQUE INDEX "main"."email"
		ON "users" (
			"email" ASC
		);`

	res, err2 := db.Exec(createTableQuery)
	util.HandleErr(err2)
	fmt.Println(res)
}

func DoLogin(email string, password string) U.User {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	var loggedUser U.User
	query := fmt.Sprintf("SELECT * FROM  users WHERE email='%v' AND password='%v'", email, password)
	rows, err := db.Query(query)
	util.HandleErr(err)
	if user, err2 := parseRows(rows); err2 != false {
		fmt.Println("Error while Parsing rows")
	} else {
		if len(user) > 0 {
			loggedUser = user[0]
		}
	}
	return loggedUser
}

func DoRegister(user U.User) bool {
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	ret := false
	insertQuery := fmt.Sprintf("INSERT INTO users(email, password, firstName, lastName, phone, address, city, userType, lat, lng) VALUES('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')", user.Email, user.Password, user.FirstName, user.LastName, user.Phone, user.Address, user.City, user.UserType, user.Lat, user.Lng)
	result, err := db.Exec(insertQuery)
	if err != nil {
		util.HandleErr(err)
	} else {
		ret = true
	}
	fmt.Println(result)
	return ret
}

func DoFetchAllUsers() []U.User {
	var users []U.User
	var err2 bool
	query := fmt.Sprintf("SELECT * FROM users")
	rows, err := db.Query(query)
	util.HandleErr(err)
	if users, err2 = parseRows(rows); err2 != false {
		fmt.Println("Error while Parsing rows")
	}

	return users
}

func DoFetchOneUser(id int) U.User {
	fetchedUser := U.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE id='%v'", id)
	rows, err := db.Query(query)
	util.HandleErr(err)
	if user, err2 := parseRows(rows); err2 != false {
		fmt.Println("Error while Parsing rows")
	} else {
		fetchedUser = user[0]
	}

	return fetchedUser
}

func parseRows(rows *sql.Rows) ([]U.User, bool) {
	error := false
	store := []U.User{}
	for rows.Next() {
		var id int64
		var email string
		var password string
		var firstName string
		var lastName string
		var phone string
		var address string
		var city string
		var userType string
		var lat float32
		var lng float32
		var status int
		var created string
		err := rows.Scan(&id, &email, &password, &firstName, &lastName, &phone, &address, &city, &userType, &lat, &lng, &status, &created)
		if err != nil {
			error = true
			util.HandleErr(err)
			return nil, error
		}
		password = "******"
		user := U.User{id, email, password, firstName, lastName, phone, address, city, userType, lat, lng, status, created}
		store = append(store, user)
	}

	return store, error
}
