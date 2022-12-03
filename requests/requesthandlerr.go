// requests provides making requests to other API's or http requests
package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
 * CreateAPIRequestAndGetResponse creates request with params encoded, converts response from request to string
 * returns error and string output of response
 */
func CreateAPIRequestAndGetResponse(value string, client *http.Client) (error, string) {
	// Create and modify HTTP request before sending
	request, err := http.NewRequest("POST", "https://www.icharts.in/includes/screener/EODScan.php", strings.NewReader(value))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	request.Header.Set("Accept-Language", "en-IN")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	req