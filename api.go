package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const url string = "http://api.reimaginebanking.com/"
const key string = "?key=a3186093b7ef7cd4d2f3d5234b9c4775"

//Customer Object
type Customer struct {
	ID        string  `json:"_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

//Address Object
type Address struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
}

//Account Object
type Account struct {
	ID         string  `json:"_id"`
	Type       string  `json:"type"`
	Nickname   string  `json:"nickname"`
	Rewards    float64 `json:"rewards"`
	Balance    float64 `json:"balance"`
	CustomerID string  `json:"customer_id"`
}

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

//SendPayment sends a payment
func SendPayment(from string, to string, amount int) {
	link := url + "accounts/" + from + "/transfers" + key
	jsonStr := `
	{
  "medium": "balance",
  "payee_id": "` + to + `",
  "amount": ` + strconv.Itoa(amount) + `,
  "transaction_date": "2017-03-25",
  "description": "Paid Using Pixiu!"
}
	`
	req, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(jsonStr)))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("Payment Status: ", resp.Status)
}

//CreateAccount Creates Capital One Account
func CreateAccount(telegramID string, telegramUsername string) {
	link := url + "accounts" + key
	jsonStr := `
{
  "first_name": "` + telegramUsername + `",
  "last_name": "` + telegramUsername + `",
  "address": {
    "street_number": "123",
    "street_name": "Main St.",
    "city": "Akron",
    "state": "OH",
    "zip": "44340"
  }
}
`

	req, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(jsonStr)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("Account Create Status: ", resp.Status)

}
