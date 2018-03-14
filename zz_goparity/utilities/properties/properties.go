package properties

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const CommentSymbol = "#"

type Properties map[string]string

func (properties *Properties) ReadProperties(fileName string) (Properties, error) {

	// dummy variable
	var outProp Properties
	outProp = Properties{"": ""}
	// check to see if the file exists
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("Error: ", fileName, " does not exist")
		return nil, err
	}

	// try to open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: Unable to open file: ", fileName)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// pass lines of a properties file into a hash map
	// for each line in the reader
	// if key exists, output error --> duplicate property exists (break out of the thing?)
	// if key does exist, insert it into hash map

	for scanner.Scan() {
		// check if line has a comment (skip it) OR if line does not contain equal sign
		// should return error if = is not present in a line (not a true properties file then)
		line := scanner.Text()
		if strings.Contains(line, CommentSymbol) || !(strings.Contains(line, "=")) {
			continue
		}
		// check to make sure that the equal sign is NOT at the beginning of the line

		fmt.Println(line)
	}

	return outProp, err
}
