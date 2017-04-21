package main

import (
	"fmt"
	"hello_world_library"
)

func main() {
	fmt.Printf("Hello golang 1.8.\n")
	fmt.Printf(stringutil.Reverse("Hello golang 1.8"))
}
