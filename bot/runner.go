package bot

import (
	"math/rand"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var BotId string

func Run() {
	err := godotenv.Load()
	errorCheck(err, "Failed to load .env file")
	// Grab token and start up bot
	token := os.Getenv("BOT_TOKEN")
	fmt.Println(token)
	bot, err := discordgo.New("Bot " + token)
	errorCheck(err, "Error authenticating bot token")
	// In this example, we only care about receiving message events.
	bot.Identify.Intents = discordgo.IntentsGuildMessages
	// Assign the bot a user
	u, err := bot.User("@me")
	errorCheck(err, "Error creating bot ID")
	// Give bot an ID
	BotId = u.ID
	// Add handler for messages
	bot.AddHandler(messageHandler)
	// Open bot
	err = bot.Open()
	errorCheck(err, "Error trying to run bot")
	// If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

// Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	command := strings.Split(message.Content, "")
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if message.Author.ID == BotId || command[0] != "!" {
		return
	}
	
	// Will likely need to cut message into parts sep=" " so that we can find the first command
	join_comm := strings.Join(command[:], "")
	split_command := strings.Split(join_comm, " ")
	// fmt.Println(split_command)
	// If check passes, try to process message
	switch split_command[0] {
	case "!ping":
		reply := replyPing(split_command)
		sendMessage(session, message, reply, "")
	case "!help":
		reply := sendHelp(split_command)
		sendMessage(session, message, reply, "")
	case "!heads":
		reply := headsTails(split_command)
		sendMessage(session, message, reply, "")
	case "!pomodor":
		PomodorQueue(split_command, session, message)
	case "!add":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong", "")
	case "!minus":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong", "")
	case "!stock":
		// idk
	case "!activity":
		reply, emoji := ActivityQueue(split_command, message.Author.Username)
		sendMessage(session, message, reply, emoji)
	}
}

// Simple function to handle the message replies
func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, reply string, emoji string) {
	if len(emoji) > 0 {
		reply = fmt.Sprint(reply, " ", emoji)
	}
	// Send reply to user message so they are tagged + ref the original message
	_, _ = s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
	fmt.Println("Message sent")
}

// Function to handle eror messages
func errorCheck(err error, message string) string {
	if err != nil {
		log.Fatal(message + err.Error())
		return message
	}
	return ""
}

// BASIC COMMANDS
// Send help link
func sendHelp(command []string) string {
	refLink := "https://github.com/sir-typesalot/JarvisBot/wiki"
	reply := "Here is a simple wiki " + refLink
	fmt.Println("Sent link to wiki")
	return reply
}
// Handle !ping status commands
func replyPing(command []string) string {
	now := time.Now()
	fmt.Println("Ping responded to at: ", now.Format("15:04:05"))
	return "Pong :ping_pong:"
}
// Heads Tails func
func headsTails(command []string) string {
	randNum := rand.Intn(50)
	if randNum % 2 == 0 {
		return "Heads :coin:"
	} else {
		return "Tails :coin:"
	}
}

