package tcct

import (
	"net"
	"log"
	"fmt"

)

// This will need to be re-written when we'll support multiple
// channel connections i guess :p
var (
// Global variable
	Tbot = NewBot()
	Name, Server, Api_key, Channel, Port = GetConfig()
)

type Bot struct {
	// The basic bot structure
	Name    string
	Server  string
	Api_key string
	Channel string
	Port    string
	Conn    net.Conn
}

func SendMessage(b *Bot, msg string) {
	// Send a message to twich channel but this is more cool.
	b.Message(msg)
}

func NewBot() *Bot {
	// Create a new bot
	return &Bot{
		Name:    Name,
		Server:  Server,
		Api_key: Api_key,
		Channel: Channel,
		Port:    Port,
		Conn:    nil,
	}
}

func (bot *Bot) Connect() {
	// Make a connection to twitch server
	var err error
	bot.Conn, err = net.Dial("tcp", bot.Server+":"+bot.Port)
	if err != nil {
		log.Panicln("Unable to connect!")
	}
}

func (bot *Bot) Message(msg string) {
	// Send a message to twitch channel
	fmt.Fprintf(bot.Conn, "PRIVMSG #"+bot.Channel+" :"+msg+"\r\n")
}