package properties

import (
	"sys"
	"os"
	"fmt"
)

type Properties map[string]


func (properties *Properties) ReadProperties(fileName string) (Properties, err){

	
	// check to see if the file exists
	if _, err := os.Stat(fileName); err != nil{
		fmt.Println("Error: ", fileName, " does not exist")
	}

	// try to open the file
	if file, err := os.Open(fileName); err != nil{
		fmt.Prinln("Error: Unable to open file: ", fileName)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// pass lines of a properties file into a hash map
	// for each line in the reader
	// if key exists, output error --> duplicate property exists (break out of the thing?)
	// if key does exist, insert it into hash map
	
}