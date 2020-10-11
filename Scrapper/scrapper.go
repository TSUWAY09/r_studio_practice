
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