package main

import (
	"mysql_db_ref/mydatabase"
	"mysql_db_ref/otherpackage"
)

func main() {

	mydatabase.Connect()

	otherpackage.Function()

}
