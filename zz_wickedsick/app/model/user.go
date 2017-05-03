package model

import (
	"log"
	"zz_wickedsick/utils/database"
	"zz_wickedsick/utils/errors"

	_ "github.com/go-sql-driver/mysql"
)

//TODO: refactor, put the structs in their own files (data classes)
type User struct {
	UserName    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	PostalCode  string
}

func (u User) AddUser() {

	dbStatement := "INSERT INTO ws_user " +
		"(date, username, password, firstname, lastname, email, phonenumber, address, postalcode) " +
		"VALUES " +
		"(CURDATE(), ?, ?, ?, ?, ?, ?, ?, ?)"

	// prepare sql statement
	stmt, err := database.DBC.Prepare(dbStatement)
	errors.HandleErr(err)

	// execute sql statement
	result, err := stmt.Exec(u.UserName, u.Password, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.Address, u.PostalCode)
	errors.HandleErr(err)

	log.Printf("user.go: ", result)

}
