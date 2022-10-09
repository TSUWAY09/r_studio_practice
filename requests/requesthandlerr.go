// requests provides making requests to other API's or http requests
package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
 * CreateAPIRequestAndGetResponse creates