package model

import (
	"log"
	"zz_wickedsick/utils/database"
	"zz_wickedsick/utils/debug"
	"zz_wickedsick/utils/errors"

	"zz_wickedsick/utils/password"

	"bytes"

	_ "github.com/go-sql-driver/mysql"
)

//TODO: refactor, put the structs in their own files (data classes)
type User struct {
	UserName       string
	Password       string
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Address        string
	PostalCode     string
	Salt           string
	HashedPassword string
}

type UserHash struct {
	UserName       string
	Password       string
	Salt           string
	HashedPassword string
}

func (u User) AddUser() {

	// set SQL queries
	dbStatement := "INSERT INTO ws_user " +
		"(date, username, password, firstname, lastname, email, phonenumber, address, postalcode) " +
		"VALUES " +
		"(CURDATE(), ?, ?, ?, ?, ?, ?, ?, ?)"

	dbStatement2 := "INSERT INTO ws_user_salt" +
		"(username, salt, saltedpassword) " +
		"VALUES " +
		"(?, ?, ?)"

	// need to insert into TWO tables
	// need to create an 'atomic' transaction for this begin and commit
	// need to prepare the txn
	txn, err := database.DBC.Begin()
	errors.HandleErr(err)

	stmt, err := txn.Prepare(dbStatement)
	errors.HandleErr(err)
	result, err := stmt.Exec(u.UserName, u.Password, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.Address, u.PostalCode)
	errors.HandleErr(err)
	id, err := result.LastInsertId()
	errors.HandleErr(err)

	// prepare second sql statemnt
	stmt2, err := txn.Prepare(dbStatement2)
	errors.HandleErr(err)
	result2, err := stmt2.Exec(u.UserName, u.Salt, u.HashedPassword)
	errors.HandleErr(err)
	id2, err := result2.LastInsertId()
	errors.HandleErr(err)

	err = txn.Commit()
	errors.HandleErr(err)

	log.Printf("user.go: User %s has been successfully added to ws_user with lastInsertId() as %d", u.UserName, id)
	log.Printf("user.go: User %s has been successfully added with ws_user_salt lastInsertId() as %d", u.UserName, id2)

	// TODO: change this so that if error does occur, return false
}

// CheckUser check if user already exists
func (u User) UserExists() bool {

	var count int
	var userExists = false
	txn, err := database.DBC.Begin()
	dbStatement := "SELECT COUNT(username) FROM ws_user where username='" + u.UserName + "'"
	debug.Log("user.go", dbStatement)
	res, err := txn.Query(dbStatement)
	errors.HandleErr(err)

	// count() is going to ALWAYS return 1 row, need to check the value of that row
	for res.Next() {
		err = res.Scan(&count)
		errors.HandleErr(err)
	}

	if count < 1 {
		debug.Log("user.go: user does not exist", u.UserName)
	} else {
		debug.Log("user.go: user already exists", u.UserName)
		userExists = true
	}

	err = txn.Commit()
	errors.HandleErr(err)

	return userExists
}

func (u User) ValidateLogin() bool {

	var dbUsername string
	var dbPassword string
	var encryptedPassword string
	var dbSalt string
	var loginResult = false

	txn, err := database.DBC.Begin()
	// can put in a check to see if there are multiple users BUT therethere are already existing checks
	// as in the username fields are UNIQUE (primary key)

	dbStatement := "SELECT * FROM ws_user_salt where username='" + u.UserName + "'"
	debug.Log("user.go", dbStatement)
	res, err := txn.Query(dbStatement)
	errors.HandleErr(err)

	for res.Next() {
		err = res.Scan(&dbUsername, &dbSalt, &dbPassword)
		errors.HandleErr(err)
	}

	encryptedPassword = password.EncryptPassword(u.Password, dbSalt)
	if (u.UserName == dbUsername) && (encryptedPassword == dbPassword) {
		loginResult = true
	}

	err = txn.Commit()
	errors.HandleErr(err)

	return loginResult

}

func GetUserDetails(userName string) (u User) {

	var dbUsername string
	var dbFirstName string
	var dbLastName string
	var dbEmail string
	var dbPhoneNumber string
	var dbAddress string
	var dbPostalCode string

	txn, err := database.DBC.Begin()
	dbStatement := "SELECT username, firstname, lastname, email, phonenumber, postalcode, address " +
		"FROM ws_user WHERE " +
		"username='" + userName + "'"
	debug.Log("user.go", dbStatement)
	res, err := txn.Query(dbStatement)
	errors.HandleErr(err)

	for res.Next() {
		debug.Log("user.go --> GetUserDetails ", "now scanning the results")
		err = res.Scan(&dbUsername, &dbFirstName, &dbLastName, &dbEmail, &dbPhoneNumber, &dbPostalCode, &dbAddress)
		errors.HandleErr(err)
	}

	u.UserName = dbUsername
	u.FirstName = dbFirstName
	u.LastName = dbLastName
	u.Email = dbEmail
	u.PhoneNumber = dbPhoneNumber
	u.PostalCode = dbPostalCode
	u.Address = dbAddress

	err = txn.Commit()
	errors.HandleErr(err)

	return u
}

func (u User) ChangeUser() {

	// if the field value is < 1 or if it's blank, assume the user didn't want to change
	// option 1
	// set fields in struct to a hash map. go through hash map and append to the SQL/UPDATE query
	// option 2
	// use reflection loop (no need for hashmap?)
	dbToUser := make(map[string]string) //make() initializes and allocates hash map

	dbToUser["username"] = u.UserName
	dbToUser["password"] = u.Password
	dbToUser["firstname"] = u.FirstName
	dbToUser["lastname"] = u.LastName
	dbToUser["email"] = u.Email
	dbToUser["phonenumber"] = u.PhoneNumber
	dbToUser["postalode"] = u.PostalCode
	dbToUser["address"] = u.Address

	// use buffer instead of string beacuse every time you concatenate a string
	// you create a NEW string in memory. Don't want that
	var newValue string
	var dbStatement bytes.Buffer
	dbStatement.WriteString("UPDATE ws_user SET ")

	// iterate through hashmap
	for key := range dbToUser {
		if len(dbToUser[key]) > 1 {
			newValue = key + "=" + dbToUser[key] + ", "
			dbStatement.WriteString(newValue)
			// is this more efficient than string concatenation?
			// dbStatement.WriteString(key)
			// dbStatement.WriteString("=")
			// dbStatement.WriteString(dbToUser[key])
			// dbStatement.WriteString(", ")
		}
	}
	log.Println(dbStatement)
}
