package main

//User struct connects Telegram and Capital One User
type User struct {
	TelegramID       string
	TelegramUsername string
	CapitalOneID     string
}

//Users is an array of the users
type Users []User

func InitUsers() {
	csinko := User{
		TelegramID:       "258303594",
		TelegramUsername: "csinko",
		CapitalOneID:     "",
	}

	kdog5 := User{
		TelegramID:       "316190324",
		TelegramUsername: "kdog5",
		CapitalOneID:     "",
	}

}
