package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type timer struct {
	user string
	timer *time.Timer
}
var timerSlice = []timer{}
var hasCanceled = []string{}

func PomodorQueue(command []string, session *discordgo.Session, message *discordgo.MessageCreate) {
	if command[1] == "cancel" {
		if status := cancelTimer(message.Author.ID); status == 200 {
			sendMessage(session, message, "You cancelled your timer ", "")
			fmt.Println("Cancelled timer")
		} else {
			sendMessage(session, message, "You don't have a timer... ", "<:mrheckles:975492543892570142>")
			fmt.Println("Non-existent timer, skipping cancellation")
		}
	} else {
		var errMsg string
		// Convert time values to numbers
		timeLength, err := strconv.ParseFloat(command[1], 32)
		// Get status of conversions
		errMsg = errorCheck(err, "Invalid time duration")
		interval, err := strconv.ParseFloat(command[2], 32)
		errMsg = errorCheck(err, "Invalid interval")
		cycles, err := strconv.Atoi(command[3])
		errMsg = errorCheck(err, "Invalid cycle")
		// Check if conversion failed
		if errMsg != "" {
			sendMessage(session, message, errMsg, "")
		} else {
			sendMessage(session, message, "Your timer has started :hourglass:", "")
			var userIdx int
			var didCancel bool
			// Loop through and create pomodor for time length
			n := 0
			for n < cycles {
				if didCancel, userIdx = userCancel(message.Author.ID); !didCancel {
					createTimer(message.Author.ID, timeLength)
					sendMessage(session, message, "Take a rest, go outside! :beach:", "")
				}
				if didCancel, userIdx = userCancel(message.Author.ID); !didCancel {
					createTimer(message.Author.ID, interval)
				}
				if n == (cycles - 1) {
					sendMessage(session, message, "Timer over, nice work! :white_check_mark:", "")
				} else {
					sendMessage(session, message, "Interval over, back to work! :keyboard:", "")
				}
				n++
			}
			if didCancel {
				hasCanceled = removeUser(hasCanceled, userIdx)
			}
		}
	}
}

func createTimer(userID string, duration float64) (int) {
	// Check if user has existing timer
	hasTimer, index := contains(timerSlice, userID)
	if hasTimer {
		log.Fatal("Has Existing Timer")
		return 401
	}
	// Multiply to convert to seconds
	minToSec := duration * 60
	timeDuration := time.Duration(minToSec)*time.Second
	// Create new timer
    new_timer := time.NewTimer(timeDuration)
	user_timer := timer{
		user: userID,
		timer: new_timer,
	}
	fmt.Println("User " + userID + " has created a timer")
	// Add struct to slice
	timerSlice = append(timerSlice, user_timer)
    <-new_timer.C
	// Get index of timer in slice
	hasTimer, index = contains(timerSlice, userID)
	// Remove slice and log
	timerSlice = removeIndex(timerSlice, index)
	return 200
}

func cancelTimer(userID string) int {
	// Check if timer exists
	checkTimer, index := contains(timerSlice, userID)
	if checkTimer {
		// Stop timer and remove from slice
		timerSlice[index].timer.Stop()
		timerSlice = removeIndex(timerSlice, index)
		return 200
	} else {
		// Return error status
		return 404
	}
}
// Check if user cancelled timer
func userCancel(userID string) (bool, int) {
	for i, users := range hasCanceled {
		if userID == users {
			return true, i
		}
	}
	return false, 0
}
func removeUser(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func contains(slice []timer, user string) (bool, int) {
	for i, timers := range slice {
		if timers.user == user {
			return true, i
		}
	}
	return false, 0
}
func removeIndex(s []timer, index int) []timer {
	return append(s[:index], s[index+1:]...)
}
