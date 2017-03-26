package main

//User struct connects Telegram and Capital One User
type User struct {
	TelegramID          string
	TelegramUsername    string
	CapitalOneID        string
	CapitalOneAccountID string
}

//Users is an array of the users
type Users []*User

func InitUsers() {
	csinko := User{
		TelegramID:          "258303594",
		TelegramUsername:    "csinko ",
		CapitalOneID:        "58d6efba1756fc834d906ad1",
		CapitalOneAccountID: "58d7ddb41756fc834d909d36",
	}

	kdog5 := User{
		TelegramID:          "316190324",
		TelegramUsername:    "kdog5 ",
		CapitalOneID:        "58d6f1571756fc834d906ad3",
		CapitalOneAccountID: "58d7de011756fc834d909d38",
	}

	memerson := User{
		TelegramID:          "304490471",
		TelegramUsername:    "memerson ",
		CapitalOneID:        "58d6f1431756fc834d906ad2",
		CapitalOneAccountID: "58d7de291756fc834d909d3c",
	}

	RepoAddUser(&csinko)
	RepoAddUser(&kdog5)
	RepoAddUser(&memerson)

}
