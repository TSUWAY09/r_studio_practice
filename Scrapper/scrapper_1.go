
package Scrapper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ScrapeContent_2() {
	client := &http.Client{}
	var data = strings.NewReader(`json=%7B%22action%22%3A%20%22advanced_search%22%2C%20%22info%22%3A%20%7B%220%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_5%22%7D%2C%20%221%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_20%22%7D%2C%20%222%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3C%22%2C%20%22cols1%22%3A%20%22bband_upper%22%7D%2C%20%223%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%224%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%225%22%3A%20%7B%22cols%22%3A%20%22atr%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2210%22%7D%2C%20%226%22%3A%20%7B%22cols%22%3A%20%22adx%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2225%22%7D%2C%20%227%22%3A%20%7B%22cols%22%3A%20%22avg_volume%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%22500000%22%7D%2C%20%228%22%3A%20%7B%22cols%22%3A%20%22p_symbol%22%2C%20%22opts%22%3A%20%22%20not%20like%20%22%2C%20%22strs%22%3A%20%22%25-%25%22%7D%2C%20%229%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22strs%22%3A%20%22%22%7D%7D%7D&_=`)

	fmt.Printf("printing data : %v", data)

	req, err := http.NewRequest("POST", "https://www.icharts.in/includes/screener/EODScan.php", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("X-Prototype-Version", "1.6.0.2")