package model

import (
	"log"
	"strconv"
	"zz_wickedsick/utils/database"
	"zz_wickedsick/utils/debug"
	"zz_wickedsick/utils/errors"

	"zz_wickedsick/utils/password"

	"database/sql"

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
	debug.Log("user.go UserExists()", dbStatement)
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
	debug.Log("user.go--> ValiidateLogin()", dbStatement)
	res, err := txn.Query(dbStatement)
	errors.HandleErr(err)

	for res.Next() {
		err = res.Scan(&dbUsername, &dbSalt, &dbPassword)
		errors.HandleErr(err)
	}

	encryptedPassword = password.EncryptPassword(u.Password, dbSalt)

	log.Println(dbUsername, dbPassword, dbSalt)
	log.Println(u.UserName, encryptedPassword)

	if (u.UserName == dbUsername) && (encryptedPassword == dbPassword) {
		debug.Log("user.go --> ValidateLogin()", "username matches and encrypted password matches")
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
	//dbToUser := make(map[string]sql.NullString) //make() initializes and allocates hash map

	var firstname sql.NullString
	var lastname sql.NullString
	var email sql.NullString
	var phonenumber sql.NullString
	var postalcode sql.NullString
	var address sql.NullString

	//this was a bitch...
	//firstname := sql.NullString{String: u.FirstName, Valid: false}
	if u.FirstName == "" {
		firstname = sql.NullString{String: u.FirstName, Valid: false}
	} else {
		firstname = sql.NullString{String: u.FirstName, Valid: true}
	}
	//lastname := sql.NullString{String: u.LastName, Valid: false}
	if u.LastName == "" {
		lastname = sql.NullString{String: u.LastName, Valid: false}
	} else {
		lastname = sql.NullString{String: u.LastName, Valid: true}
	}
	//email := sql.NullString{String: u.Email, Valid: false}
	if u.Email == "" {
		email = sql.NullString{String: u.Email, Valid: false}
	} else {
		email = sql.NullString{String: u.Email, Valid: true}
	}
	//phonenumber := sql.NullString{String: u.PhoneNumber, Valid: false}
	if u.PhoneNumber == "" {
		phonenumber = sql.NullString{String: u.PhoneNumber, Valid: false}
	} else {
		phonenumber = sql.NullString{String: u.PhoneNumber, Valid: true}
	}
	//postalcode := sql.NullString{String: u.PostalCode, Valid: false}
	if u.PostalCode == "" {
		postalcode = sql.NullString{String: u.PostalCode, Valid: false}
	} else {
		postalcode = sql.NullString{String: u.PostalCode, Valid: true}
	}
	//address := sql.NullString{String: u.Address, Valid: false}
	if u.Address == "" {
		address = sql.NullString{String: u.Address, Valid: false}
	} else {
		address = sql.NullString{String: u.Address, Valid: true}
	}

	txn, err := database.DBC.Begin()

	// use ifnull() mysql function
	dbStatement := "UPDATE ws_user SET firstname=IFNULL(?, firstname), " +
		"lastname=IFNULL(?, lastname), " +
		"email=IFNULL(?, email), " +
		"phonenumber=IFNULL(?, phonenumber), " +
		"postalcode=IFNULL(?, postalcode), " +
		"address=IFNULL(?, address) " +
		"WHERE username='" + u.UserName + "'"

	debug.Log("user.go --> ChangeUser", dbStatement)
	stmt, err := txn.Prepare(dbStatement)
	errors.HandleErr(err)

	res, err := stmt.Exec(firstname, lastname, email, phonenumber, postalcode, address)
	errors.HandleErr(err)

	numRowsAffected, err := res.RowsAffected()
	errors.HandleErr(err)
	if numRowsAffected < 1 {
		log.Println("unable to update" + u.UserName + " affected rows: " + strconv.FormatInt(numRowsAffected, 10))
	} else if numRowsAffected > 1 {
		panic("too many rows updated, panic")
	} else {
		debug.Log("user.go --> ChangeUser()", "update successful with rows affected: "+strconv.FormatInt(numRowsAffected, 10))
	}

	err = txn.Commit()
	errors.HandleErr(err)

} //TODO: refactor - do these functions need to be struct methods or shuld they be separate from the struct

//  ChangePassword TODO: set this to return bool?
func (u User) ChangePassword() {

	// generate a new salt
	newSalt := password.GenerateSalt()

	// hash the salt and password together
	newHashedPassword := password.EncryptPassword(u.Password, newSalt)

	txn, err := database.DBC.Begin()
	dbStatement := "UPDATE ws_user_salt SET salt='" + newSalt + "', " + "saltedpassword='" + newHashedPassword + "' " + "WHERE username='" + u.UserName + "'"
	debug.Log("user.go -->ChangePassword()", dbStatement)
	res, err := txn.Exec(dbStatement)
	errors.HandleErr(err)

	numRowsAffected, err := res.RowsAffected()
	errors.HandleErr(err)

	if numRowsAffected < 1 {
		log.Println("unable to update" + u.UserName + " affected rows: " + strconv.FormatInt(numRowsAffected, 10))
	} else if numRowsAffected > 1 {

		panic("too many rows updated, panic")
	} else {
		debug.Log("user.go --> ChangeUser()", "update successful with rows affected: "+strconv.FormatInt(numRowsAffected, 10))
	}

	dbStatement = "UPDATE ws_user SET password='" + u.Password + "' " + "WHERE username='" + u.UserName + "'"
	debug.Log("user.go -->ChangePassword()", dbStatement)
	res, err = txn.Exec(dbStatement)
	errors.HandleErr(err)

	numRowsAffected, err = res.RowsAffected()
	errors.HandleErr(err)

	if numRowsAffected < 1 {
		log.Println("unable to update" + u.UserName + " affected rows: " + strconv.FormatInt(numRowsAffected, 10))
	} else if numRowsAffected > 1 {
		// txn rollback
		panic("too many rows updated, panic")
	} else {
		debug.Log("user.go --> ChangeUser()", "update successful with rows affected: "+strconv.FormatInt(numRowsAffected, 10))
	}

	err = txn.Commit()
	errors.HandleErr(err)

}

func (u User) DeleteUser() {

	txn, err := database.DBC.Begin()
	dbStatement := "Delete from ws_user where username='" + u.UserName + "'"
	debug.Log("user.go -->DeleteUser()", dbStatement)
	res, err := txn.Exec(dbStatement)
	errors.HandleErr(err)

	numRowsAffected, err := res.RowsAffected()
	errors.HandleErr(err)
	if numRowsAffected < 1 {
		log.Println("unable to delete" + u.UserName + " affected rows: " + strconv.FormatInt(numRowsAffected, 10))
	} else if numRowsAffected > 1 {
		// txn rollbakc
		panic("too many rows updated, panic")
	} else {
		debug.Log("user.go --> DeleteUser()", "update successful with rows affected: "+strconv.FormatInt(numRowsAffected, 10))
	}

	dbStatement = "Delete from ws_user_salt where username='" + u.UserName + "'"
	debug.Log("user.go -->DeleteUser()", dbStatement)
	res, err = txn.Exec(dbStatement)
	errors.HandleErr(err)

	numRowsAffected, err = res.RowsAffected()
	errors.HandleErr(err)
	if numRowsAffected < 1 {
		log.Println("unable to delete" + u.UserName + " affected rows: " + strconv.FormatInt(numRowsAffected, 10))
	} else if numRowsAffected > 1 {
		// txn rollbakc
		panic("too many rows updated, panic")
	} else {
		debug.Log("user.go --> DeleteUser()", "update successful with rows affected: "+strconv.FormatInt(numRowsAffected, 10))
	}

	err = txn.Commit()
	errors.HandleErr(err)

}
