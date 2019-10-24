package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/andersfylling/disgord"
)

type config struct {
	ClientID    uint
	Secret      string
	Permissions uint
	Token       string
}

var cum config
var discord *disgord.Client

func init() {
	initializeConfig()
}

func main() {
	startDc()
}

func initializeConfig() {
	c, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Could not find config.json - terminating app.")
	}
	defer c.Close()

	cm, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatal("Could not read config.json - terminating app")
	}
	err = json.Unmarshal(cm, &cum)
	if err != nil {
		log.Fatal("Could not parse config.json - terminating app")
	}
}

func startDc() {
	discord = disgord.New(&disgord.Config{
		BotToken: cum.Token,
	})
	defer discord.StayConnectedUntilInterrupted()

	dcOnMessage(discord)

	discord.Ready(func() {
		fmt.Println("Bot running")
		fmt.Println(disgord.LibraryInfo())
	})
}

func dcOnMessage(dc *disgord.Client) {
	dc.On(disgord.EvtMessageCreate, func(s disgord.Session, evt *disgord.MessageCreate) {
		msg := evt.Message
		msgIntoArr := strings.Fields(msg.Content)
		aut := evt.Message.Author
		autTag := aut.Username
		cur, err := s.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
			return
		}
		if aut.Bot {
			return
		}
		if len(msgIntoArr) < 1 {
			return
		}

		if strings.Contains(msgIntoArr[0], cur.ID.String()) {
			r := fmt.Sprintf("<@%v> why are you taggin me?", aut.ID)
			msg.Reply(s, r)
			return
		}

		switch strings.ToLower(msgIntoArr[0]) {
		case "!help":

			h := fmt.Sprintf(
				`Publicly available commands are as follows:

'!faq' - Frequently asked questions
'!tobase64 message' - Encodes message to base64
'!frombase64 message' - Decodes message from base64
'!ping' - Pong.`)
			msg.Reply(s, h)

		case "!faq":
			m := fmt.Sprintf(
				`
**FAQ - Frequently asked questions:**

Q1) How long does this course take to finish?
A1) This really depends on each person. Some may complete in 2 weeks, some may complete in 2 months, it's up to each person. But in order to get traction in programming a person should spend at least 10 hours a week coding.

Q2) Will I be able to land a job after this course? Is there people who got hired after this course?
A2) This course alone will not be sufficient to get a job without connections. A developer has to grow portfolio of personal projects in modern tech world to prove he/she is capable. There are many of us that got hired after completing course but most of this time learning additional stuff was necessary. For more questions please refer to <#445233695633440778>

Q3) Can I ask a coding question?
A3) Of course, please use <#445233618466373661> for your code related questions and provide as much as details as possible, if you are sharing code wrap it with triple backticks.

Q4) What are good sources to learn x?
A4) Please try "!suggestfree" and "!suggestcheap"

Q5) I'm too confused, I don't know what to learn next.
A5) This is perfectly normal, web development is a gigantic field. Please refer to this <https://github.com/kamranahmedse/developer-roadmap>. This roadmap might scare you at first, but remember. You do not need to know everything, just things you need for your goals.
`)
			msg.Reply(s, m)
			m = fmt.Sprintf(
				`
Q6) 
`)
			msg.Reply(s, m)

		case "!suggestfree":
			if len(msgIntoArr) < 2 {
				m := `Please provide one of the following topics for me to suggest courses about
- Example: 
	"!suggestfree react"

Possible options: css, node, javascript`
				msg.Reply(s, m)
				return
			}
			switch strings.ToLower(msgIntoArr[1]) {

			case "css":
				m := `Free CSS Courses: 
freeCodeCamps Responsive Web Design Certification <https://www.freecodecamp.org/learn>
LET'S GET GRIDDY WITH IT (CSS GRID with Wes Bos) <https://cssgrid.io/>
WHAT THE FLEXBOX?! (Flexbox with Wes Bos) - <https://flexbox.io/>`
				msg.Reply(s, m)

			case "node":
				m := `Free NodeJS Courses:
Code with Node: Learn by Doing - <https://www.devsprout.io/code-with-node>`
				msg.Reply(s, m)

			case "javascript":
				m := `Free JavaScript Courses:
freeCodeCamps Algorithms and Data Structures Certification <https://www.freecodecamp.org/learn>
JavaScript30 (ES6 and DOM manipulation with Wes Bos) <https://javascript30.com/>`
				msg.Reply(s, m)
			}

		case "!suggestcheap":
			if len(msgIntoArr) < 2 {
				m := `Please provide one of the following topics for me to suggest courses about
- Example: 
	"!suggestcheap react"

Possible options: react, css, javascript, typescript, reactnative, webdesign, php, go
`
				// I moved the node option to !suggestfree as Ian has made it free on devsprout.io
				msg.Reply(s, m)
				return
			}
			switch strings.ToLower(msgIntoArr[1]) {
			case "react":
				m := `Cheap ReactJS Courses:
The Modern React Bootcamp (Hooks, Context, NextJS, Router) - <https://www.udemy.com/course/modern-react-bootcamp/> 
React - The Complete Guide (incl Hooks, React Router, Redux) - <https://www.udemy.com/share/101WayB0IfcVxXRn4=/>
(Advanced) React Testing with Jest and Enzyme - <https://www.udemy.com/share/101ZdQB0IfcVxXRn4=/>`
				// I figured Colts course should be the top most option.
				msg.Reply(s, m)

			case "css":
				m := `Cheap CSS Courses:
Advanced CSS and Sass: Flexbox, Grid, Animations and More! - <https://www.udemy.com/share/101WmqB0IfcVxXRn4=/>`
				msg.Reply(s, m)

			case "javascript":
				m := `Cheap JavaScript Courses:
JavaScript: Understanding the Weird Parts - <https://www.udemy.com/course/understand-javascript/>
The Complete JavaScript Course 2019: Build Real Projects! - <https://www.udemy.com/course/the-complete-javascript-course/>`
				msg.Reply(s, m)
			case "typescript":
				m := `Cheap TypeScript Courses:
Typescript: The Complete Developer's Guide - <https://www.udemy.com/share/101X9oB0IfcVxXRn4=/>`
				msg.Reply(s, m)
			case "reactnative":
				m := `Cheap React Native Courses:
React Native - The Practical Guide: <https://www.udemy.com/share/101WwKB0IfcVxXRn4=/>`
				msg.Reply(s, m)
			case "webdesign":
				m := `Cheap Web Design Courses:
Adobe Photoshop CC - Web Design, Responsive Design & UI - <https://www.udemy.com/share/101W4aB0IfcVxXRn4=/>`
				msg.Reply(s, m)
			case "php":
				m := `Cheap PHP Courses:
PHP for Beginners - Become a PHP Master - CMS Project - <https://www.udemy.com/share/101X5QB0IfcVxXRn4=/>`
				msg.Reply(s, m)
			case "go":
				m := `Cheap Go (Golang) Courses:
Learn How To Code: Google's Go (golang) Programming Language - <https://www.udemy.com/share/101r9AB0IfcVxXRn4=/>
(Requires basic Go knowledge) Web Development w/ Google’s Go (golang) Programming Language - <https://www.udemy.com/share/1022eCB0IfcVxXRn4=/>`
				msg.Reply(s, m)
			}

		case "!debug":
			fmt.Println(evt.Message.Content)
		case "!tobase64":
			if len(msgIntoArr) < 2 {
				msg.Reply(s, "Encode what? Please provide additional message to encode.")
				return
			}
			n := strings.Join(msgIntoArr[1:], " ")
			str := base64.StdEncoding.EncodeToString([]byte(n))
			msg.Reply(s, str)

		case "!frombase64":
			if len(msgIntoArr) < 2 {
				msg.Reply(s, "Decode what? Please provide additional message to decode.")
				return
			}
			n := strings.Join(msgIntoArr[1:], " ")
			str, err := base64.StdEncoding.DecodeString(n)
			if err != nil {
				msg.Reply(s, "There has been an error decoding the string:\n", err)
				return
			}
			msg.Reply(s, string(str))

		case "!ping":
			msg.Reply(s, "pong")

		case "!actedit":
			if autTag != "yittoo#7826" {
				msg.Reply(s, "You aint Yit")
				return
			}
			if len(msgIntoArr) < 2 {
				msg.Reply(s, "You need to provide a status name after actEdit")
				return
			}
			act := disgord.NewActivity()
			n := strings.Join(msgIntoArr[1:], " ")
			act.Name = n

			upd := disgord.UpdateStatusCommand{
				Game: act,
			}

			err = s.UpdateStatus(&upd)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}
