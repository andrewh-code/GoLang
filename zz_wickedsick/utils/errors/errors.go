package errors

import (
	"log"
)

func HandleErr(err error) {

	if err != nil {
		// use log.Fatal, outputs the error to the log and then calls os.exit(1)
		log.Fatal(err)

	}
}
