package dao

import (
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
	createTableQuery := `create table if not exists users(
		id integer not null primary key autoincrement,
		username text,
		email text,
		password text
	)`

	res, err2 := db.Exec(createTableQuery)
	util.HandleErr(err2)
	fmt.Println(res)
}

func DoLogin(username string, password string) U.User {

	var loggedUser U.User
	query := fmt.Sprintf("select * from  users where username='%v' and password='%v'", username, password)
	rows, err := db.Query(query)
	for rows.Next() {
		var username string
		var email string
		var id int64
		var password string
		err1 := rows.Scan(&id, &username, &email, &password)
		util.HandleErr(err1)
		loggedUser = U.User{id, username, email, password}

	}
	util.HandleErr(err)
	return loggedUser
}

func DoRegister(username string, email string, password string) bool {
	ret := false
	insertQuery := fmt.Sprintf("insert into users(username, email, password) values('%v', '%v', '%v')", username, email, password)
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
	users := []U.User{}
	query := fmt.Sprintf("select * from  users")
	rows, err := db.Query(query)
	util.HandleErr(err)
	for rows.Next() {
		var username string
		var email string
		var id int64
		var password string
		err1 := rows.Scan(&id, &username, &email, &password)
		util.HandleErr(err1)
		users = append(users, U.User{id, username, email, password})
	}
	return users
}

func DoFetchOneUser(id int) U.User {
	fetchedUser := U.User{}
	query := fmt.Sprintf("select * from  users where id='%v'", id)
	rows, err := db.Query(query)
	util.HandleErr(err)
	for rows.Next() {
		var username string
		var email string
		var id int64
		var password string
		err1 := rows.Scan(&id, &username, &email, &password)
		util.HandleErr(err1)
		fetchedUser = U.User{id, username, email, password}
	}

	return fetchedUser
}
