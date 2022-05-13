package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func StocksQueue(command []string) {
	
	switch command[1] {
	case "market":
		fmt.Println("market status")
	case "symbol":
		fmt.Println("symbol data")
	}
    
}

func marketStatus() {
	// GET request to API
	resp, err := http.Get("https://api.polygon.io/v1/marketstatus/now?apiKey=HFL_1uqtU8zFVeJhl09fcWTE9sSNRcSw")
	errorCheck(err, "Failed to GET resource")
	// Delay closing of resp Body
    defer resp.Body.Close()
	// Convert to var
	body, err := ioutil.ReadAll(resp.Body)
	errorCheck(err, "Failed to translate body text")
	// Print out
    fmt.Println(string(body))
}
