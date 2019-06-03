package main

import (
	"fmt"

	"time"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	TELEGRAM_APITOKEN = "813476983:AAHH-jfBwzqWgXUA_jGQc4YY8UBJgyTxYBA"

	exTheme1 = "🎯 Chess & Triceps"
	exTheme2 = "🎯 Leg & Shoulders"
	exTheme3 = "🎯 Back & Biceps"

	// 📌\"\" - \n
	exDay1 = "Execirses for today:\n\n📌\"Жим лежа\" - 4х10\n📌\"Жим хаммер\" - 4х10\n📌\"Разводка гантель\" - 4х12\n📌\"Разводка на кросовере\" - 4х12\n📌\"Отжимания от скамьи\" - 4х10\n📌\"Французский жим на трицепс\" - 4х10\n📌\"Француская косичка\" - 4х12\n"
	exDay2 = "Execirses for today:\n\n📌\"Присид\" - 4х10\n📌\"Выпады\" - 3х20\n📌\"Разгибания ног\" - 4х10\n📌\"Икры\" - 4х20\n📌\"Жим Арнольда\" - 4х10\n📌\"Разводка гантель на плечи\" - 4х15\n📌\"Тяга грифа к подбородку\" - 4х12\n"
	exDay3 = "Execirses for today:\n\n📌\"Гиперэкстензия\" - 4х12\n📌\"Тяга верх. блока\" - 4х12\n📌\"Горизонт тяга\" - 4х12\n📌\"Тяга гантель на спину\" - 4х10\n📌\"21\" - 3х2\n📌\"Кривой гриф обр. хват\" - 4х12\n📌\"Гантель на бицепс\" - 4х12\n📌\"Пресс\" - 3х30\n"
)

var days = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(exTheme1),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(exTheme2),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(exTheme3),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("❎ Exit"),
	),
)
var menu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⚙ Menu"),
	),
)

type Conversation struct {
	User *tgbotapi.User
}

func NewConversation(User *tgbotapi.User) *Conversation {
	return &Conversation{
		User: User,
	}
}

var conversations = map[int]*Conversation{}

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Error("Error of updating")
	}

	for update := range updates {
		// if update.Message == nil {
		// 	continue
		// }

		// if !update.Message.IsCommand() {
		// 	continue
		// }

		User := update.Message.From
		UserName := User.FirstName
		UserID := User.ID
		now := time.Now()
		timenow := "⏱ " + now.Format("02.01.2006 15:04")
		seperator := "\n---------------------------------------------\n"

		//
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		conv := conversations[UserID]

		if conv != nil {
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "stop" {
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.Text = "✅ Conversation was successfully deleted.\n See you " + UserName + ".\n\nTo starting new conversation type: /start" + seperator + "Author: @xbaysal11"
					delete(conversations, UserID)
				} else if cmdText == "menu" {
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.Text = `Click to "Menu" button and choose exercise`
					msg.ReplyMarkup = menu
				} else {
					msg.Text = "❌ Conversation already exist!\n\nTo cancel type: /stop"
				}
			} else {
				if update.Message.Text == menu.Keyboard[0][0].Text {
					msg.Text = "🏋‍♂ Good luck  " + UserName + "!"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = days
				} else if update.Message.Text == days.Keyboard[0][0].Text {
					fmt.Println(now)
					msg.Text = exTheme1 + seperator + timenow + seperator + exDay1
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = menu
				} else if update.Message.Text == days.Keyboard[1][0].Text {
					msg.Text = exTheme2 + seperator + timenow + seperator + exDay2
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = menu
				} else if update.Message.Text == days.Keyboard[2][0].Text {
					msg.Text = exTheme3 + seperator + timenow + seperator + exDay3
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = menu
				} else if update.Message.Text == days.Keyboard[3][0].Text {
					msg.Text = "✅ Conversation was successfully deleted.\n See you " + UserName + ".\n\nTo starting new conversation type: /start" + seperator + "Author: @xbaysal11"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					delete(conversations, UserID)
				} else {
					// other messages
					msg.Text = "I'm bot 🤖.\nI don't know this command.\nYou can start new conversation by type: /start"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				}
			}
		} else {
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "start" {
					conversations[UserID] = NewConversation(User)
					msg.Text = "✋ Hello, " + UserName + ".\n\n" + `Click to "Menu" button and choose exercise`
					msg.ReplyMarkup = menu
				} else {
					msg.Text = "I'm bot 🤖.\nI don't know this command.\nYou can start new conversation by type: /start"
				}
			}
		}
		// send
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
