package main

import (
	"database/sql"
	"fmt"

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
	// sqlQuery := "SELECT * FROM " + DB_TABLE
	// sqlDelete := "DELETE FROM user WHERE username='andrew'"

	stmtInsert, err := dbc.Exec(sqlInsert)
	if err != nil {
		panic(err)
	}
	fmt.Println("stmtInsert", stmtInsert)

	// //query result
	// stmtQuery, err := dbc.Query(sqlQuery)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("stmtQuery is: ", stmtQuery)

	// //delete
	// stmtDelete, err := dbc.Exec(sqlDelete)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("stmtDelete is: ", stmtDelete)

	//close the database connection
	dbc.Close()
}
