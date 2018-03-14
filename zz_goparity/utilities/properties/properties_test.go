package properties

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {

	// declare propeties struct variable
	var out Properties
	var err error

	test, err := out.ReadProperties("config.properties")
	if err != nil {
		t.Fail()
	}
	fmt.Println(test)
}
