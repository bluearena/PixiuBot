package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

func testapi() {
	url := "http://api.reimaginebanking.com/accounts?key=fab8d98dc85bc29afbe6b9915f27a0e7"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func SentPayment(from string, to string, amount float64) {
	link := url + "accounts/" + from + "/transfers" + key
	jsonStr := `
	{
  "medium": "balance",
  "payee_id": "` + to + `",
  "amount": ` + strconv.FormatFloat(amount, 'E', -1, 64) + `,
  "transaction_date": "2017-03-25",
  "description": "Paid Using Pixiu!"
}
	`
	_, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println(err)
	}

}
