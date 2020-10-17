
// Scrapper provides scrapping API for scrapping stock related data with bullish and bearish stocks
package Scrapper

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"stockScanner/communication"
	"stockScanner/fileio"
	"stockScanner/filters"
	"stockScanner/requests"
	"stockScanner/types"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/html"
)

/*
 * ScrapeContent takes command line arguments for options
 * -t stock option type
 * -s simple technical details only
 */
func ScrapeContent(args []string, options []string) {
	selections := filters.ConstructArgsMap(args)
	log.Println(selections)
	index := 0
	// Create HTTP client with timeout
	client := &http.Client{}

	bearishStocksScanner := `{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": "<", "cols1": "ema_20"}, "1": {"cols": "last_close", "opts": "<", "cols1": "sma_50"}, "2": {"cols": "dmi_plus", "opts": "<", "cols1": "dmi_minus"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "6": {"cols": "atr", "opts": ">", "strs": "10"}, "7": {"cols": "adx", "opts": ">", "strs": "25"}, "8": {"cols": "rsi", "opts": ">", "strs": "30"}, "9": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}}}`
	bullishStocksScanner := `{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": ">", "cols1": "ema_5"}, "1": {"cols": "last_close", "opts": ">", "cols1": "ema_20"}, "2": {"cols": "last_close", "opts": "<", "cols1": "bband_upper"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "atr", "opts": ">", "strs": "10"}, "6": {"cols": "adx", "opts": ">", "strs": "25"}, "7": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "8": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}, "9": {"cols": "", "opts": " like ", "strs": ""}}}`
	params := url.Values{}
	stockScannerType := filters.GetStockScannerType(args)
	log.Printf("stock scanner type received is %d \n", stockScannerType)

	// sets required stockScannerType in params for URL encoding from command line argument
	if types.BULLISH == stockScannerType {
		log.Printf("creating bullish stock scanner object")
		params.Set("json", bullishStocksScanner)
	} else {
		log.Printf("creating bearish stock scanner object")
		params.Set("json", bearishStocksScanner)
	}
	value := params.Encode()
	/*log.Println("received encoded string is : ", value)*/

	err, htmlOutput := requests.CreateAPIRequestAndGetResponse(value, client)
	if err != nil {
		log.Panic("error while creating API request or while reading response")
	}
	//log.Println("****************************************************")
	//log.Printf("response : %+v \n", htmlOutput)
	log.Println("****************************************************")

	htmlTokens := html.NewTokenizer(strings.NewReader(htmlOutput))

	// outputs bullish or bearish based on the stockScannerType
	if types.BULLISH == stockScannerType {
		log.Println("*******************BULLISH STOCKS*******************")
	} else {
		log.Println("*******************BEARISH STOCKS*******************")
	}
	log.Println("****************************************************")

	var stocksData [][]string // holds data for both rows and columns when written to csv file
	var rowData []string      // holds data for each row
	i := 0
	for i < 1 {

		tt := htmlTokens.Next()