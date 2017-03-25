package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("348577595:AAHHm0zC10iUEp6l5o2RP-pJ45Hl9f9DOnU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	response := "Well heck that didn't work maybe you should try again!"

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("MESSAGE TEXT[%s]", update.Message.Text)

		messageArray := strings.Split(update.Message.Text, " ")
		fmt.Printf("MESSAGE ARRAY: %v", messageArray)

		//The command - put in error checking to make sure this is a valid command
		slashCommand := messageArray[0]
		//fmt.Printf(slashCommand)
		switch slashCommand {
		case "/pay":
			fmt.Println("You hit the pay command")
			response = "You hit the pay command"
			//Call pay function
		case "/donger":
			fmt.Println("You hit the donger command")
			response = "You hit the donger command"
			//call the donger function
		}

		//Check if the second thing has an at name
		//Then check to see if the username is on the object or whatever

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func pay() {

	//DO a bunch of calls

}
