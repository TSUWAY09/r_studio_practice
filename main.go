package main

import (
	"log"
	"os"
	"stockScanner/Scrapper"
	"stockScanner/filters"
)

/*
 * command line application for generating bullish/bearish stocks
 */
func main() {
	// Read command 