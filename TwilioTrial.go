package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	// Set account keys & information

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	// Create possible message bodies
	quotes := [5]string{"It Really Be like that sometimes", "Hit the yeet", "Hit or Miss", "Heck Yeah", "Big Chungus Hours"}

	// Set up rand
	rand.Seed(time.Now().Unix())

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", ("+" + os.Getenv("PHONE_NUMBER")))//verified number
	msgData.Set("From", ("+" + os.Getenv("TWILIO_NUMBER")))//twilio account number
	msgData.Set("Body", quotes[rand.Intn(len(quotes))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
