package main

import (
	"mysql_db_ref/mydatabase"
)

func main() {

	mydatabase.Connect()

	mydatabase.Query()

}
