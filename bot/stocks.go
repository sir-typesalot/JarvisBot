package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Polygon.io URL for market data
const Polygon_URL = "https://api.polygon.io/v1";

func StocksQueue(command []string) (response string, emoji string) {
	
	switch command[1] {
	case "status":
		response, emoji = marketStatus()
	case "symbol":
		response, emoji = "", ""
	}
	return
    
}

func marketStatus() (string, string) {
	// Craft response
	type Response struct {
		Market 		string `json:"market"`
		EarlyHours  bool `json:"earlyHours"`
		AfterHours 	bool `json:"afterHours"`
		ServerTime  string `json:"serverTime"`
		Exchanges struct {
			NYSE string `json:"nyse"`
			NASDAQ string `json:"nasdaq"`
			OTC string `json:"otc"`
		} `json:"exchanges"`
		Currencies struct {
			Fx string `json:"fx"`
			Crypto string `json:"crypto"`
		} `json:"currencies"`
	}
	
	var errList []error
	var emoji string

	err := godotenv.Load()
	errList = errorCheck(err, "Failed to load .env file", errList)
	// Grab token for API
	apiToken := os.Getenv("POLY_API_TOKEN")
	// Craft URL
	url := fmt.Sprintf("%s/marketstatus/now?apiKey=%s", Polygon_URL, apiToken)

	// GET request to API
	resp, err := http.Get(url)
	errList = errorCheck(err, "Failed to GET resource", errList)
	// Delay closing of resp Body
    defer resp.Body.Close()
	// Read response in
	body, err := ioutil.ReadAll(resp.Body)
	errList = errorCheck(err, "Failed to read body", errList)
	
	stockInfo := Response{}
	err = json.Unmarshal(body, &stockInfo)
	errList = errorCheck(err, "Could not read body", errList)

	if len(errList) > 0 {
		emoji = "<:cat_cry:975383207996456980>"
		return "We're having some issues getting data from the API ", emoji
	}
	// TODO: Pretty this output up
	reply := ""
	if stockInfo.Market == "open" { 
		emoji = ":white_check_mark:"
	} else {
		emoji = ":octagonal_sign:"
	}
	s := fmt.Sprintf("Market Status - %s %s\nNASDAQ - %s\nNYSE - %s\nForex - %s\nCrypto - %s", 
		stockInfo.Market, emoji, stockInfo.Exchanges.NASDAQ, stockInfo.Exchanges.NYSE,
		stockInfo.Currencies.Fx, stockInfo.Currencies.Crypto)
	reply += s
	return reply, ""
}
