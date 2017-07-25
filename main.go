package main

import (
	"flag"
	"fmt"
	"syscall"
	"unicode"

	"strings"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Token for bot")
	flag.Parse()
}

func main() {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	session.AddHandler(respond)

	err = session.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running. Use CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

func respond(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!pig-latin") {
		fmt.Println("Message: ", m.Content)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> %s", m.Author.ID, pigLatin(m.Content)))
	}

}

func pigLatin(s string) string {
	var runes = []rune(strings.TrimLeft(strings.ToLower(s), "!pig-latin"))
	var head rune
	var pig []rune

	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]) {
			if head == 0 {
				switch runes[i] {
				case 'a', 'e', 'i', 'o', 'u':
					pig = append(pig, runes[i])
				}
				head = runes[i]
			} else {
				pig = append(pig, runes[i])
			}
		} else {
			switch head {
			case 'a', 'e', 'i', 'o', 'u':
				pig = append(pig, 'w', 'a', 'y')
			case 0:
			default:
				pig = append(pig, head, 'a', 'y')
			}
			head = 0
			pig = append(pig, runes[i])
		}
	}

	return string(pig)
}
