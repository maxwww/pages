package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"strconv"
	"strings"
)

func NewBot(bot *tgbotapi.BotAPI, messages Messages) *Bot {
	return &Bot{
		Bot:      bot,
		Messages: messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.Bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message != nil {
			firstPage := 0
			lastPage := 0
			re1, _ := regexp.Compile(`(\d+)\s(\d+)`)
			res1 := re1.FindAllStringSubmatch(update.Message.Text, -1)
			if len(res1) > 0 {
				firstPage, _ = strconv.Atoi(res1[0][1])
				lastPage, _ = strconv.Atoi(res1[0][2])
				delta := (lastPage - firstPage + 1) % 4
				if delta > 0 {
					lastPage = lastPage + (4 - delta)
				}
			} else {
				re2, _ := regexp.Compile(`(\d+)\+(\d+)`)
				res2 := re2.FindAllStringSubmatch(update.Message.Text, -1)
				if len(res2) > 0 {
					firstPage, _ = strconv.Atoi(res2[0][1])
					totalLists, _ := strconv.Atoi(res2[0][2])
					lastPage = totalLists*4 + (firstPage - 1)
				}
			}
			if firstPage == 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.Messages.Responses.Description)
				b.Bot.Send(msg)
			}
			var firstSet []string
			var secondSet []string
			iterations := (lastPage - firstPage + 1) / 4
			for i := 0; i < iterations; i++ {
				firstSet = append(firstSet, strconv.Itoa(lastPage-(2*i)), strconv.Itoa(firstPage+(2*i)))
				secondSet = append(secondSet, strconv.Itoa((firstPage+1)+(2*i)), strconv.Itoa((lastPage-1)-(2*i)))
			}

			firstLine := strings.Join(firstSet, ", ")
			secondLine := strings.Join(secondSet, ", ")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, firstLine+"\n"+secondLine)
			b.Bot.Send(msg)
		}
	}

	return nil
}
