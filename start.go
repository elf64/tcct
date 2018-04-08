package main

import (
	"awesomeProject2/tcct"
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
	"bufio"
	"net/textproto"
	"strings"
)


func main() {
	// GUI STUFF
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(tcct.Layout)

	// TWITCH CONNECTION STUFF
	tcct.Tbot.Connect()
	fmt.Fprintf(tcct.Tbot.Conn, "PASS %s\r\n", tcct.Tbot.Api_key)
	fmt.Fprintf(tcct.Tbot.Conn, "NICK %s\r\n", tcct.Tbot.Name)
	fmt.Fprintf(tcct.Tbot.Conn, "JOIN #%s\r\n", tcct.Tbot.Channel)

	reader := bufio.NewReader(tcct.Tbot.Conn)
	tp := textproto.NewReader(reader)

	go func() {
		// Main loop for the twitch chat
		for {
			line, err := tp.ReadLine()
			if err != nil {
				break
			}
			if strings.Contains(line, "PING") {
				// Calm your apples!! PING-PONG works
				p := strings.Split(line, "PING ")
				fmt.Fprintf(tcct.Tbot.Conn, "PONG %s\r\n", p[1])
			}
			if strings.Contains(line, "tmi.twitch.tv PRIVMSG #"+tcct.Tbot.Channel) {
				msg := strings.Split(line, "tmi.twitch.tv PRIVMSG #"+tcct.Tbot.Channel)[1]
				user := strings.Split(line, "!")[0]
				user = strings.Replace(user, ":", "", -1)
				if user == tcct.Tbot.Name {
					continue
				}
				if strings.Contains(msg, tcct.Tbot.Name) {
					tcct.UpdateChat(g, "\x1b[0;0m"+user+" ->\x1b[0;31m"+msg+"\x1b[0;0m")
				} else {
					tcct.UpdateChat(g, "\x1b[0;0m"+user+"\x1b[0;0m ->"+msg)
				}
				//if msg == " :!test" {
				//	Tbot.Message("TEST")
				//}
			}
		}
	}()

	// Keybindings for GUI and GUI Mainloop
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, tcct.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, tcct.UpdateView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}