
// types provide all data types used in stockScanner
package types

import (
	"gopkg.in/yaml.v2"
)

// used for determining whether a stock selected is Bearish or Bullish
const (
	DUMMY = iota
	BULLISH
	BEARISH
)

type Params struct {
	Key   string
	Value Value
}

type filter string

type Value struct {
	Action string