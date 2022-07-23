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
	// Read command line arguments
	argsWithoutProg := os.Args[1:]
	allOptions := filters.GetCommandLineOptions(argsWithoutProg)
	log.Printf("