package model

// set global database connection

// // global variables
// var DBC *sql.DB
// var err error

// //var db config.DBStruct = config.RetrieveDBConfiguration(FILE_NAME)
// var dbConnectCommand string = "root:password@tcp(localhost:3333)/godb_sandbox"

// // //connect to database
// func connect(dbConnectCommand string) {
// 	DBC, err = sql.Open("mysql", dbConnectCommand)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestConnect(t *testing.T) {
// 	//open database connection
// 	DBC, err = sql.Open("mysql", dbConnectCommand)
// 	if err != nil {
// 		t.Error("cannot connect to database", dbConnectCommand)
// 	}
// }

// func TestUserExists(t *testing.T) {

// 	DBC, err = sql.Open("mysql", dbConnectCommand)
// 	var isPass bool
// 	var u User
// 	u.UserName = "test"

// 	isPass, err = u.UserExists()
// 	if err != nil || isPass == false {
// 		t.Error(err, isPass)
// 	}
// }
