package main

import (
	"fmt"

	"github.com/ikarabanov/json"
)

type EmailConfiguration struct {
	Sender         string
	Receiver       string
	SMTPUserID     string
	SMTPPassword   string
	SMTPServer     string
	SMTPPortNumber string
}

func main() {
	var l *EmailConfiguration
	err := readJson.LoadJSON("EmailConf.json", &l)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Pass: ", l)

	err = readJson.WriteJSON("test.json", l)
	if err != nil {
		fmt.Println(err)
	}

}
