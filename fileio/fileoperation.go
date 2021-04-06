// fileio provides different operations related to file operations
package fileio

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stockScanner/types"
	"time"
)

/* CreateCSVFile creates csv file and returns file pointer with current time as file name in the format 20060102150405
 * GO follows below constants for formatting of date
	const (
    stdLongMonth      = "January"
    stdMonth          = "Jan"
    stdNumMonth       = "1"
    stdZeroMonth      = "01"
    stdLongWeekDay    = "Monday"
    stdWeekDay        = "Mon"
    stdDay            = "2"
    stdUnderDay       = "_2"
    stdZeroDay        = "02"
    stdHour           = "15"
    stdHour12         = "3"
    stdZeroHour12     = "03"
    stdMinute         = "4"
    stdZeroMinute     = "04"
    stdSecond         = "5"
    stdZeroSecond     = "05"
    stdLongYear       = "2006"
    stdYear           = "06"
    stdPM             = "PM"
    stdpm             = "pm"
    stdTZ             = "MST"
    stdISO8601TZ      = "Z0700"  // prints Z for UTC
    stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
    stdNumTZ          = "-0700"  // always numeric
    stdNumShortTZ     = "-07