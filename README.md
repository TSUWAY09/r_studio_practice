
# stockScanner
stockScanner application for tracking strong bullish and bearish stocks based on monthly, 21day, 6 month, bollinder bands

## stock variant
`go run main.go -t bullish -o file` should provide a csv file with all bullish stocks
`go run main.go -t bearish -o file` should provide a csv file with all bearish stocks

## simple option
`go run main.go -t bullish -o file -s` options s provides simple output instead of many technical values
if option s is not mentioned all technical values will be updated in csv sheet
 
 ## table view in command line
 `go run main.go -t bullish -o table -s` writes all data to a table format in command line