package properties

import (
	"bufio"
	"fmt"
	"os"
)

type Properties map[string]string

func (properties *Properties) ReadProperties(fileName string) (Properties, error) {

	// dummy variable
	var outProp Properties
	outProp = Properties{"": ""}
	// check to see if the file exists
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("Error: ", fileName, " does not exist")
	}

	// try to open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: Unable to open file: ", fileName)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// pass lines of a properties file into a hash map
	// for each line in the reader
	// if key exists, output error --> duplicate property exists (break out of the thing?)
	// if key does exist, insert it into hash map

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return outProp, err
}
