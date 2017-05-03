package debug

import (
	"log"
)

/*
can use a variadic function
func function(x ..int){} --> accepts x amount of int arguments (in golang, it's not method/function overlaoding)
but to the developer, it is
limitation to this is that the argumetns ALL have to be of the same type
*/

/*
can use interfaces
*/

var debugFlag bool

func Init(flag bool) {
	debugFlag = flag
}

func Log(fileName string, input string) {
	//global variable
	if debugFlag == true {
		log.Printf(fileName + ": " + input)
	}
}

/*
ackage main
func main() {
       var t bool
       z:=&t
       *z = true
}
*/
