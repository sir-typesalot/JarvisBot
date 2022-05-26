package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// TODO: For now, just trimmed the date strings - should be a better solution
// Format the date strings
func ActivityQueue(command []string, author string) (response string, emoji string) {

	// Load dotenv
	err := godotenv.Load()
	errorCheck(err, "Failed to load .env file")
	// Grab URL
	BASE_URL := os.Getenv("EXERCISE_API_URL")
	// TODO: improvise the functions here, maybe create a struct with url, comm, author?
	if len(command) < 2 {
		return "Not enough args", "<:laughingtom:975383179601010718>"
	}
	switch command[1] {
	case "user-info":
		response = getUserInfo(BASE_URL, command, author)
		emoji = ""
	case "create-user":
		response, emoji = createUser(BASE_URL, command, author)
	case "log":
		response, emoji = logActivity(BASE_URL, command, author)
	case "user-stats":
		response, emoji = getUserStats(BASE_URL, command, author)
	case "scoreboard":
		response, emoji = getScoreboard(BASE_URL, command)
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
	errorCheck(err, "Failed to read body")
	// Unmarshal result, check for err
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errorCheck(err, "Could not read body")

	// Log output and return data as string
	fmt.Printf("The username is %s", user_info.Data.Username)
	reply := fmt.Sprintf(
		"**Username**: %s\n**UserID**: %d\n**Registered**: %s",
		user_info.Data.Username, user_info.Data.ID, user_info.Data.CreateTime)
	return reply
}

// NOTE: Will likely need to chenge this soon, if the format of endpoint in API changes
func getUserStats(url string, command []string, author string) (string, string) {
	// Create data structs
	type DataRow struct {
		Activity string `json:"total_activity"`
		Records  int `json:"total_records"`
	}
	type UserInfo struct {
		Data []DataRow `json:"data"`
		User string    `json:"user"`
	}
	var username string
	var errMsg string
	emoji := ""

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
	errMsg = errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errMsg = errorCheck(err, "Failed to read body")
	
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errMsg = errorCheck(err, "Could not read body")

	if errMsg != "" {
		emoji = "<:cat_cry:975383207996456980>"
		return errMsg, emoji
	}
	// Format data and return as string
	reply := ""
	for _, row := range user_info.Data {
		s := fmt.Sprintf("User **%s** has clocked in a total of **%s** minutes over %d records", 
		user_info.User, row.Activity, row.Records)
		reply += s
	}
	return reply, emoji
}

func getScoreboard(url string, command []string) (string, string) {
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

	var errMsg string
	var timeRange string
	var endpoint string
	
	emoji := ""

	if len(command) > 2 {
		timeRange = command[2]
	} else {
		timeRange = "all-time"
	}
	
	isValid := isValidRange(timeRange)

	if isValid {
		endpoint = "/activity/scoreboard?timeframe="+timeRange
	} else {
		emoji = "<:cat_cry:975383207996456980>"
		return "Invalid time range", emoji
	}
	
	// Craft endpoint
	completeURL := fmt.Sprint(url, endpoint)
	resp, err := http.Get(completeURL)
	// TODO: Instead of reassigning same value, maybe create slice or errors
	// And then run a check through the list?
	errMsg = errorCheck(err, "Failed to GET API")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errMsg = errorCheck(err, "Failed to read body")
	
	user_info := UserInfo{}
	err = json.Unmarshal(body, &user_info)
	errMsg = errorCheck(err, "Could not read body")

	if errMsg != "" {
		emoji = "<:cat_cry:975383207996456980>"
		return errMsg, emoji
	}

	var reply string
	if len(user_info.Data) > 0 {
		fmt.Printf("The date is %s\n", user_info.Data[0].Date)
		for _, row := range user_info.Data {

			datetime_split := strings.Split(row.Date, " ")
			date := datetime_split[:4]
			s := fmt.Sprintf("**User**: %s **Date**: %s **Time**: %s\n\n", row.Username, strings.Join(date, " "), row.Activity)
			reply += s
		}
		return reply, emoji
	} else {
		return "No Data for today, get up and do something", "<:risitas:975382207625584640>"
	}
	
}

func createUser(url string, command []string, author string) (string, string) {
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
	errorCheck(err, "Failed to read body")

	bodyJson := RespBody{}
	err = json.Unmarshal(body, &bodyJson)
	errorCheck(err, "Could not read body")
	// Return response based on status codes
	fmt.Printf("User created with code %d\n", bodyJson.Status)
	switch bodyJson.Status {
	case 200:
		s := fmt.Sprintf("Successfully created user @%s ", username)
		return s, "<:woo_baby:975382482050514955>"
	case 300:
		s := fmt.Sprintf("User already exists with name @%s ", username)
		return s, "<:risitas:975382207625584640>"
	case 500:
		s := "Encountered an error trying to create user, let's try again "
		return s, "<:cat_cry:975383207996456980>"
	default:
		return "Ran into some problems trying to create user ", "<:cat_cry:975383207996456980>"
	}
}

func logActivity(url string, command []string, author string) (string, string) {

	type RespBody struct {
		Status string `json:"status"`
		Code   int    `json:"code"`
	}
	// Check if minutes are in command
	if len(command) < 3 {
		return "You forgot to add the minutes ", "<:laughingtom:975383179601010718>"
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
	errorCheck(err, "Failed to read body")

	// bodyString := string(body)
	// fmt.Print(bodyString)

	bodyJson := RespBody{}
	err = json.Unmarshal(body, &bodyJson)
	errorCheck(err, "Could not read body")
	// Return response based on status code
	fmt.Printf("Logged activity for user %s\n", author)
	if bodyJson.Code == 200 {
		return "Your records are safe with me ", "<:leo_cheers:975383282755715112>"
	} else {
		return "I didn't catch that, let's try again", "<:cat_cry:975383207996456980>"
	}
}

func isValidRange(str string) bool {
	validRanges := [3]string {"all-time", "week", "day"}
	for _, v := range validRanges {
		if v == str {
			return true
		}
	}

	return false
}
