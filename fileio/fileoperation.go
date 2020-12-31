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
 * GO follo