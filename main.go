package main
import (
	"fmt"
    "strconv"
	"gopkg.in/telegram-bot-api.v4"
	"log"
    //"net/http"
    //"io/ioutil"
	"strings"
    "errors"
)

func main() {

	//Initialize the users in the chat
	InitUsers()
	fmt.Println("Initialized the Users")

	bot, err := tgbotapi.NewBotAPI("348577595:AAHHm0zC10iUEp6l5o2RP-pJ45Hl9f9DOnU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	response := "Default response"

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("MESSAGE TEXT[%s]", update.Message.Text)

		messageArray := strings.Split(update.Message.Text, " ")
		fmt.Println("MESSAGE ARRAY: %v", messageArray)

        if len(messageArray) == 0 {
            response = "Zero arguments detected"
        }

		//The command - put in error checking to make sure this is a valid command
		slashCommand := messageArray[0]
		//fmt.Printf(slashCommand)
		switch slashCommand {

		case "/pay":
			fmt.Println("You hit the pay command")
            response = pay(messageArray)

		case "/donger":
			fmt.Println("You hit the donger command")
			response = "You hit the donger command"
			//call the donger function
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func pay(messageArray []string) string {
    returnString := "default"
    
    if(len(messageArray) < 3) {
		return "Error: There are less than three arguments"
    }
	
	//Get the username, and check that its valid
	atUser := messageArray[1]
	valid := ValidUser(atUser)
	if valid == false {
		return "Error: Username not found"
	}

	//Get the dollar amount and check that its valid
    amountToSend, err := strconv.Atoi(messageArray[2])
    err = validPrice(amountToSend)
    if err != nil {
		returnString = "Error: Invalid Price"
		return returnString
	}

	returnString = "We made a payment"

    return returnString
}

func validPrice(price int) error{

    if (price < 0) {
        return errors.New("invalid price: negative value")
    }

    return nil;
}

func ValidUser(user string) bool {
	valid := false
	user = user[1:len(user)]
	for i := 0; i < len(users); i++ {
		if user == users[i].TelegramUsername {
			valid = true
		}
	}
	return valid
}
