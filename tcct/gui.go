package tcct

import (
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
	"strings"
	"time"
)

const (
	VERSION = "0.2.2"
)

func GetBuffView(g *gocui.Gui) (string, error) {
	// Returns string, nil
	// String is the buffer from view Input
	v, err := g.View("input")
	if err != nil {
		log.Panicln(err)
	}
	buff := v.Buffer()
	return buff, nil
}

func UpdateChat(g *gocui.Gui, msg string) {
	// This function is used in the twitch loop to update
	// the GUI with the new messages from twitch chat.
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("chat")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Fprintln(v, msg)
		return nil
	})
}

func UpdateView(g *gocui.Gui, v *gocui.View) error {
	// Because we're cool we're updating the chat/input view BUUUUUTT
	// we're also making a new connection to twitch and sending the
	// message!!! can you believe it???!?!
	//^^^ That was removed but i'm too lazy to delete the comment
	// Here we update the chat view
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("chat")
		if err != nil {
			log.Panicln(err)
		}
		//v.Clear()
		buff, _ := GetBuffView(g)
		fmt.Fprint(v, "You -> " + buff)
		return nil
	})

	// Seems like :/ we need to give the GUI enough time to
	// print out message to the screen.. i guess 0.1 should do
	time.Sleep(100 * time.Millisecond)
	msg, _ := GetBuffView(g)
	Tbot.Message(msg)

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("input")
		if err != nil {
			log.Panicln(err)
		}
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})

	return nil
}

func SetView(g *gocui.Gui, name string) (*gocui.View, error) {
	// Set the view
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func SpamLine(c string,times int) string {
	// Spam a chararacter
	// Here we make a string of size times and then we just append
	// the character to it and return the slice as a string.
	first := make([]string,times,times)
	for i := 0; i < times; i++ {
		first[i] = c
	}
	final := strings.Join(first, "")
	return final
}

func Layout(g *gocui.Gui) error {
	// Here we create the views and set the pos of them
	maxX, maxY := g.Size()
	if v, err := g.SetView("chat", 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
		v.Wrap = true
		spam := SpamLine("-", 50)
		fmt.Fprintln(v, "\x1b[0;32mWelcome!\x1b[0;35m -"+
			" Version:"+VERSION+" - Twitch Chat Client made in Go\x1b[0;0m"+
			"\n\x1b[0;31m"+spam+"\x1b[0;0m\n"+
			"Currently connected to: "+Server+
			"\nCurrent Channel: "+Channel+
			"\nCurrent Name: "+Name+
			"\n\x1b[0;31m"+spam+"\x1b[0;0m\n")
	}

	if v, err := g.SetView("input", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = true
		if _, err = SetView(g, "input"); err != nil {
			return err
		}
	}

	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	// Quit function for GUI
	return gocui.ErrQuit
}




