//Pointers
package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sandbox/subpackage"
)

// CustomStruct blah blah blah
type CustomStruct struct {
	X string
	Y int
}

// don't return anything
func pointerStuff() {

	var i = 52
	j := 72 // inference

	pI := &i
	pJ := &j

	fmt.Println("pointer location is: %d", pI)
	fmt.Println("pointer value is: %d", *pI)
	fmt.Println("pointer location is: %d", pJ)
	fmt.Println("pointer value is: %d", *pJ)

	*pI = 40
	fmt.Println("%d", *pI)

	fmt.Println(CustomStruct{"Hello", 10})

	// Pointer to struct
	customStructValue := CustomStruct{"hello world", 100}
	pointertoCustomSTructValue := &customStructValue

	fmt.Println(customStructValue)
	fmt.Println(&pointertoCustomSTructValue)
	fmt.Println(*pointertoCustomSTructValue)

	//change struct values inside pointer
	pointertoCustomSTructValue.X = "goodbye world"
	fmt.Println(customStructValue)
	fmt.Println(*pointertoCustomSTructValue)
}

func hash() {

	var strMsg = "this is the original string before the sha1"
	var hash = sha1.New()
	hash.Write([]byte(strMsg))
	var sha1Hash = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	//hex.EncodetoString(hash.Sum(nil))

	fmt.Println(strMsg)
	fmt.Println(sha1Hash)
}

func main() {
	//pointerStuff()
	//hash()
	fmt.Println(returnNumber())

	subpackage.SubPackageHello()
	subpackage.SubPackageHelloAgain()
}
