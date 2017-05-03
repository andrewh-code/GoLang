package otherpackage

import (
	"fmt"
	"log"
	"mysql_db_ref/mydatabase"
	"strconv"
)

func Function() {

	log.Println("now selecting from table user")
	rows, err := mydatabase.DBC.Query("select * from user")
	if err != nil {
		panic(err)
	}

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
}
