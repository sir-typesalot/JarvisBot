package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	
	var errMsg string
	var emoji string

	// GET request to API
	resp, err := http.Get("https://api.polygon.io/v1/marketstatus/now?apiKey=HFL_1uqtU8zFVeJhl09fcWTE9sSNRcSw")
	errMsg = errorCheck(err, "Failed to GET resource")
	// Delay closing of resp Body
    defer resp.Body.Close()
	// Read response in
	body, err := ioutil.ReadAll(resp.Body)
	errMsg = errorCheck(err, "Failed to read body")
	
	stockInfo := Response{}
	err = json.Unmarshal(body, &stockInfo)
	errMsg = errorCheck(err, "Could not read body")

	if errMsg != "" {
		emoji = "<:cat_cry:975383207996456980>"
		return errMsg, emoji
	}
	// TODO: Pretty this output up
	reply := ""
	s := fmt.Sprintf("Market Status\nMarket - %s\nNASDAQ - %s\nNYSE - %s", 
		stockInfo.Market, stockInfo.Exchanges.NASDAQ, stockInfo.Exchanges.NYSE)
	reply += s
	return reply, emoji
}
