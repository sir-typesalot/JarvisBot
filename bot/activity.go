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
	// TODO: improvise the functions here, maybe create a struct with url, comm, author?
	if len(command) < 2 {
		return "Not enough args :laughingtom:"
	}
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

func getUserInfo(url string, command []string, author string) string {
	// Create struct to match payload
	type UserInfo struct {
		Data struct {
			CreateTime string `json:"created_datetime"`
			Groupname  string `json:"group_name"`
			ID         int    `json:"id"`
			Username   string `json:"username"`
		} `json:"data"`
	}
	var username string
	// Check if username has been provided, if not, use author
	// TODO: Move this to a function (that struct is looking kinda good rn)
	if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}
	// Craft specific endpoint 
	endpoint := "/activity/user/"
	completeURL := fmt.Sprint(url, endpoint, username)
	// GET from API and check for err
	resp, err := http.Get(completeURL)
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Unmarshal result, check for err
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	// Log output and return data as string
	fmt.Printf("The username is %s", user_info.Data.Username)
	reply := fmt.Sprintf(
		"Username: %s\nUserID: %d\nRegistered: %s",
		user_info.Data.Username, user_info.Data.ID, user_info.Data.CreateTime)
	return reply
}

// NOTE: Will likely need to chenge this soon, if the format of endpoint in API changes
func getUserStats(url string, command []string, author string) string {
	// Create data structs
	type DataRow struct {
		Activity string `json:"activity"`
		Date     string `json:"date"`
	}
	type UserInfo struct {
		Data []DataRow `json:"data"`
		User string    `json:"user"`
	}
	var username string
	// TODO: Really need to move this to a new function
	if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}
	// Craft endpoint and add uesrname to it
	endpoint := "/activity/user/"
	completeURL := fmt.Sprint(url, endpoint, username, "/stats")
	resp, err := http.Get(completeURL)
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	// Format data and return as string
	reply := "User: " + user_info.User + "\n"
	for _, row := range user_info.Data {
		s := fmt.Sprintf("Date: %s \tActive Minutes: %s\n", row.Date, row.Activity)
		reply += s
	}
	return reply
}

func getScoreboard(url string, command []string) string {
	// Create data structs
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
	// Craft endpoint
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
	// Creaft response 
	type RespBody struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
		Status int `json:"status"`
	}
	var username string

	if len(command) > 2 {
		username = command[2]
	} else {
		username = author
	}

	// Create json data to POST to API
	values := map[string]string{"username": username}
	json_data, err := json.Marshal(values)
	errorCheck(err, "Error parsing values to JSON")
	// Craft endpoint and POST
	endpoint := "/activity/user/create"
	completeURL := fmt.Sprint(url, endpoint)
	resp, err := http.Post(completeURL, "application/json", bytes.NewBuffer(json_data))
	errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodyJson := RespBody{}
	err = json.Unmarshal(body, &bodyJson)
	errorCheck(err, "Could not read body")
	// Return response based on status codes
	fmt.Printf("User created with code %d\n", bodyJson.Status)
	switch bodyJson.Status {
	case 200:
		s := fmt.Sprintf("Successfully created user @%s :woo_baby:", username)
		return s
	case 300:
		s := fmt.Sprintf("User already exists with name @%s :risitas:", username)
		return s
	case 500:
		s := "Encountered an error trying to create user, let's try again :cat_cry:"
		return s
	default:
		return "Ran into some problems trying to create user :cat_cry:"
	}
}

func logActivity(url string, command []string, author string) string {

	type RespBody struct {
		Status string `json:"status"`
		Code   int    `json:"code"`
	}
	// Check if minutes are in command
	if len(command) < 3 {
		return "You forgot to add the minutes :laughingtom:"
	}
	// Convert to int
	minutes_str := command[2]
	minutes, err := strconv.Atoi(minutes_str)
	errorCheck(err, "Cannot convert str to int")
	// Create json to POST
	values := map[string]int{"minutes": minutes}
	json_data, err := json.Marshal(values)
	errorCheck(err, "Error parsing values to JSON")
	// Craft endpoint and send
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
	// Return response based on status code
	fmt.Printf("Logged activity for user %s\n", author)
	if bodyJson.Code == 200 {
		return "Your records are safe with me :leo_cheers:"
	} else {
		return "I didn't catch that, let's try again"
	}
}
