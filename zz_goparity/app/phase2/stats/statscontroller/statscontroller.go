package statscontroller

import (
	"fmt"
	"net/http"
)

func HelloWorldStats(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from stats")
}
