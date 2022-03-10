
// filters provides utilities for handling various utility operations
package filters

import (
	"log"
	"sort"
	"stockScanner/types"
	"strings"

	"golang.org/x/net/html"
)

/*
 *	ReplaceUnnecessaryHtmlData replaces unnecessary html data left over after extracting required data from html page
 *	uses strings.ReplaceAll() to replace old content with "" string	in given input parameter eg
 */
func ReplaceUnnecessaryHtmlData(eg html.Token) string {
	actualString := strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", "")
	actualString = strings.ReplaceAll(actualString, `<a href="\&#34;javascript:" toajaxtableeditor('order_by_changed',="" new="" array('`, "")
	actualString = strings.ReplaceAll(actualString, `','desc'));\"="">`, "")
	actualString = strings.ReplaceAll(actualString, `','asc'));\"="">`, "")
	return actualString
}

/* GetStockScannerType reads arguments when run and if option mentioned is -t and value is BEARISH/BULLISH
* accordingly types.BEARISH/ types.BULLISH will be returned if none is present or no option is selected then
* types.BULLISH will be selected by default
 */
func GetStockScannerType(args []string) int {
	//checks if given argument contains t then reads next stock option
	if strings.Contains(args[0], "t") {
		log.Printf("ouptut stock scanner type is selected")
		if strings.EqualFold(args[1], "Bearish") {
			return types.BEARISH
		}
	}
	return types.BULLISH
}

/*
 * GetCommandLineOptions takes all command line arguments as input
 * returns all the options chosen
 */