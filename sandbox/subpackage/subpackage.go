package subpackage

import "fmt"

// remember to CAPITALIZE the first letter of the functions for public access
// lowercase indicates the identifier is private and only accessed by the package it was declared

func SubPackageHello() {
	fmt.Println("hello world from subpackage")
}

func SubPackageHelloAgain() {
	fmt.Println("hello world from subpackage again")
}
