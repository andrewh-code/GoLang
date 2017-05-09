package errorstuff

import (
	"log"
)

func HandleErr(err error) {

	if err != nil {
		// use log.Fatal, outputs the error to the log and then calls os.exit(1)
		log.Fatal(err)

	}
}

func DBErr(err error) {
	if err != nil {
		log.Printf(err.Error())
		log.Printf("now exiting db function")
	}

}
