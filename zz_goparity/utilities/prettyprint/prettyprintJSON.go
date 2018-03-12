package prettyprint

import (
	"encoding/json"
	"fmt"
)

// PrettyPrintJSON - literally a wrapper for json.MarshalIndent with 4 spaces as the default tabbing
func PrettyPrintJSON(v interface{}) []byte {

	jsonOut, err := json.MarshalIndent(v, "", "    ")
	// TODO: don't necessarily need the error?
	// can either set the logging here in the prety print or the logging in the fucntion that calls it
	if err != nil {
		// TODO: replace with log
		fmt.Println("error:", err)
	}

	return jsonOut
}
