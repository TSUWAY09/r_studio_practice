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
	request, err := http.NewRequest("POST", "ht