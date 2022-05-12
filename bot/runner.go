package bot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var BotId string
var bot *discordgo.Session

func Run() {
	err := godotenv.Load()
	errorCheck(err, "Failed to load .env file")
	// Grab token and start up bot
	token := os.Getenv("BOT_TOKEN")
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
	// If check passes, try to process message
	switch message.Content {
	case "!ping":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	case "!help":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	case "!heads":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	case "!pomodor":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	case "!add":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	case "!minus":
		// Send reply to user so they are tagged
		sendMessage(session, message, "pong")
	}
}

// Simple function to handle the message replies
func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, reply string) {
	_, _ = s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
	fmt.Print("Message sent")
}

// Function to handle eror messages
func errorCheck(err error, message string) {
	if err != nil {
		log.Fatal(message + err.Error())
	}
}

// // function baseComms(message) {
//   if (message.content == '!ping') {
//     // Reply to message and log response time
//     message.reply('Pong');
//     var now = new Date();
//     var timeRn = now.getHours() + ":" + now.getMinutes() + ":" + now.getSeconds();
//     // Log time
//     console.log(`Ping responded to at ${timeRn}`);
//   } else if (message.content == '!help') {
//     // Send user link to wiki page
//     refLink = 'https://github.com/sir-typesalot/JarvisBot/wiki'
//     message.reply(`Here is a simple reference sheet ${refLink}`);
//     console.log('Sent Link to Wiki');
//   } else {
//     console.log('Message not understood');
//   }
// };
