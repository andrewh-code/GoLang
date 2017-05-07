package debug

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/* http response fields
Method:
URL:
Protocol:
Host:
Accept-Encoding: --> acceptable encodings (gzip, deflate, etc)
Accept-Language: --> acecptable languages
Cache-Control: --> specify directives that must be obeyed by all caching mechanisms along ther request-response chain
User-Agent: --> application identifier (identifies itself, application type, OS, software vendor as a string), transmitted as a header field user-agent
Referer: --> address of previous web page from which link to current request was followed
Connection: --> control options for current connection and list of hop-by-hop response fields
Pragma: --> Implementation-specific fields that may have various effects anywhere along the request-response chain.
Accept: --> content types acceptable for response (ie text/plain)

ex)
2017/05/02 00:23:49 ===============BEGIN REQUEST=================================
Method: GET
 url: /
 protocol: HTTP/1.1
 host: localhost:8080
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36
Accept-Language: en-US,en;q=0.8
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8
Accept-Encoding: gzip, deflate, sdch, br
================END REQUEST================================
*/

func FormatRequest(r *http.Request) {
	// Create return string
	var request []string

	// break point
	request = append(request, "\n")
	request = append(request, "===============BEGIN REQUEST=================================")

	// format the method, url, protocol, and hostURL
	url := fmt.Sprintf("Method: %v \n url: %v \n protocol: %v \n host: %v", r.Method, r.URL, r.Proto, r.Host)
	request = append(request, url)

	// get the rest of the header information
	for headerName, headerDetails := range r.Header {
		for _, h := range headerDetails {
			request = append(request, fmt.Sprintf("%v: %v ", headerName, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n\nPOST Request Payload Information: \n")
		request = append(request, r.Form.Encode())
	}

	request = append(request, "================END REQUEST================================\n\n")
	// Return the request as a string
	log.Printf(strings.Join(request, "\n"))
}
