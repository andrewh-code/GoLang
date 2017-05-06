package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// constants
const (
	DB_USER  = "root"
	DB_NAME  = "GODB_SANDBOX"
	PASSWORD = "password"
	DB_HOST  = "tcp(localhost:3333)"

	DB_TABLE = "user"
)

type User struct {
	UserName sql.NullString
}

func main() {

	dataSourceName := DB_USER + ":" + PASSWORD + "@" + DB_HOST + "/" + DB_NAME
	fmt.Println(dataSourceName)

	//open database connection
	dbc, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	//db queries

	//db insert
	sqlInsert := "INSERT INTO user (username, password, email, date) values ('andrew', 'password', 'andrew@email.com', '2017-01-01')"
	fmt.Println(sqlInsert)
	sqlQuery := "SELECT * FROM " + DB_TABLE
	sqlDelete := "DELETE FROM user WHERE username='andrew'"

	stmtInsert, err := dbc.Prepare("INSERT INTO user (username, password, email, date) values (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	dbcResult, err := stmtInsert.Exec("andre", "password123", "andre@gmail.com", "2017-01-01")
	if err != nil {
		panic(err) //turn this into its own function
	}
	id, err := dbcResult.LastInsertId()

	fmt.Println("dbcResult: ", id)

	//query result
	rows, err := dbc.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("stmtQuery is: ", rows)

	// go through the results
	for rows.Next() {
		var uid int
		var username string
		var password string
		var email string
		var date string

		err = rows.Scan(&uid, &username, &password, &email, &date)
		fmt.Println(strconv.Itoa(uid) + "\t" + username + "\t" + password + "\t" + email + "\t" + date)
	}

	//delete
	stmtDelete, err := dbc.Exec(sqlDelete)
	if err != nil {
		panic(err)
	}
	fmt.Println("stmtDelete is: ", stmtDelete)

	//sql null types
	var u User
	var temp string
	temp = "false"
	u.UserName = sql.NullString{String: temp, Valid: false} // if true, then pass in "" or '' to the database. If FALSE, pass in NULL to the database
	sqlUpdate := "update user set password=ifnull(?, password) where username='andre'"
	stmtUpdate, err := dbc.Prepare(sqlUpdate)
	if err != nil {
		panic(err)
	}
	result, err := stmtUpdate.Exec(u.UserName)
	if err != nil {
		panic(err)
	}
	log.Println(result)

	rows, err = dbc.Query("select * from user")
	for rows.Next() {
		var uid int
		var username string
		var password string
		var email string
		var date string

		err = rows.Scan(&uid, &username, &password, &email, &date)
		fmt.Println(strconv.Itoa(uid) + "\t" + username + "\t" + password + "\t" + email + "\t" + date)
	}
	//close the database connection
	dbc.Close()
}
