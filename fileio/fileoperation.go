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
    stdDay     