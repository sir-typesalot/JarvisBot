package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// TODO: Need to parse the dates better
// Format the date strings
func ActivityQueue(command []string, author string) (response string) {

	// Load dotenv
	err := godotenv.Load()
	errorCheck(err, "Failed to load .env file")
	// Grab URL
	BASE_URL := os.Getenv("EXERCISE_API_URL")

	switch command[1] {
	case "user-info":
		response = getUserInfo(BASE_URL, command, author)
	case "create-user":
		response = createUser(BASE_URL, command, author)
	case "log":
		response = logActivity(BASE_URL, command, author)
	case "user-stats":
		response = getUserStats(BASE_URL, command, author)
	case "scoreboard":
		response = getScoreboard(BASE_URL, command)
	}
	return
}

// Need a list of command
// prefix - !activity
// user-info <username>, gives info on user specified
// create-user <username?>, creates user from username, if no name, then message author
// log <time>, will log time with message author
// user-stats <username>, return stats of user
// scoreboard

func getUserInfo(url string, command []string, author string) string {

	type UserInfo struct {
		Data struct {
			CreateTime string `json:"created_datetime"`
			Groupname  string `json:"group_name"`
			ID         int    `json:"id"`
			Username   string `json:"username"`
		} `json:"data"`
	}
	var username string

	if len(command) < 2 {
		return "Not enough args"
	} else if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}
	endpoint := "/activity/user/"
	completeURL := fmt.Sprint(url, endpoint, username)
	resp, err := http.Get(completeURL)
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Unmarshal result
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	fmt.Printf("The username is %s", user_info.Data.Username)
	reply := fmt.Sprintf(
		"Username: %s\nUserID: %d\nRegistered: %s",
		user_info.Data.Username, user_info.Data.ID, user_info.Data.CreateTime)
	return reply
}

func getUserStats(url string, command []string, author string) string {
	type DataRow struct {
		Activity string `json:"activity"`
		Date     string `json:"date"`
	}
	type UserInfo struct {
		Data []DataRow `json:"data"`
		User string    `json:"user"`
	}
	var username string

	if len(command) < 2 {
		return "Not enough args"
	} else if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}
	endpoint := "/activity/user/"
	completeURL := fmt.Sprint(url, endpoint, username, "/stats")
	resp, err := http.Get(completeURL)
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	reply := "User: " + user_info.User + "\n"
	fmt.Printf("The username is %s\n", user_info.User)
	fmt.Printf("The date is %s\n", user_info.Data[0].Date)
	for _, row := range user_info.Data {
		s := fmt.Sprintf("Date: %s \tActive Minutes: %s\n", row.Date, row.Activity)
		reply += s
	}
	return reply
}

func getScoreboard(url string, command []string) string {

	type DataRow struct {
		Activity string `json:"activity"`
		Date     string `json:"date"`
		Username string `json:"username"`
	}
	type UserInfo struct {
		Data     []DataRow `json:"data"`
		Metadata struct {
			Count      int `json:"count"`
			TotalUsers int `json:"num_users"`
		} `json:"metadata"`
	}

	endpoint := "/activity/scoreboard"
	completeURL := fmt.Sprint(url, endpoint)
	resp, err := http.Get(completeURL)
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	var reply string
	fmt.Printf("The date is %s\n", user_info.Data[0].Date)
	for _, row := range user_info.Data {
		s := fmt.Sprintf(":id: %s :date: %s Active Minutes: %s\n\n", row.Username, row.Date, row.Activity)
		reply += s
	}
	return reply
}

func createUser(url string, command []string, author string) string {

	type RespBody struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
		Status int `json:"status"`
	}
	var username string

	if len(command) < 2 {
		return "Not enough args"
	} else if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}

	values := map[string]string{"username": username}
	json_data, err := json.Marshal(values)
	errorCheck(err, "Error parsing values to JSON")

	endpoint := "/activity/user/create"
	completeURL := fmt.Sprint(url, endpoint)
	resp, err := http.Post(completeURL, "application/json", bytes.NewBuffer(json_data))
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodyJson := RespBody{}
	err = json.Unmarshal(body, &bodyJson)
	errorCheck(err, "Could not read body")

	fmt.Printf("User created with code %d\n", bodyJson.Status)
	if bodyJson.Status == 200 {
		s := fmt.Sprintf("Successfully created user @%s", username)
		return s
	} else {
		s := "Encountered an error trying to create user, let's try again"
		return s
	}
}

func logActivity(url string, command []string, author string) string {

	type RespBody struct {
		Status string `json:"status"`
		Code   int    `json:"code"`
	}
	if len(command) < 3 {
		return "You forgot to add the minutes"
	}
	minutes_str := command[2]
	minutes, err := strconv.Atoi(minutes_str)
	errorCheck(err, "Cannot convert str to int")

	values := map[string]int{"minutes": minutes}
	json_data, err := json.Marshal(values)
	errorCheck(err, "Error parsing values to JSON")

	endpoint := "/activity/user/"
	completeURL := fmt.Sprint(url, endpoint, author, "/log-activity")
	resp, err := http.Post(completeURL, "application/json", bytes.NewBuffer(json_data))
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// bodyString := string(body)
	// fmt.Print(bodyString)

	bodyJson := RespBody{}
	err = json.Unmarshal(body, &bodyJson)
	errorCheck(err, "Could not read body")

	fmt.Printf("Logged activity for user %s\n", author)
	if bodyJson.Code == 200 {
		return "Your records are safe with me"
	} else {
		return "I didn't catch that, let's try again"
	}
}
