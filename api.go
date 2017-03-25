package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {
	url := "http://api.reimaginebanking.com/accounts?key=fab8d98dc85bc29afbe6b9915f27a0e7"
	//url := "https://google.com"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
