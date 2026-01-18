package main

import (
	"fmt"
	sendnotification "mail/sendNotification"
)

func main() {

	sliceAllMessage := &sendnotification.AllMessage{
		SendMessage: []sendnotification.Message{},
	}


	newMessage := sendnotification.Message{Title: "Hola",Description: "Description"}
	sliceAllMessage.SendMessage = append(sliceAllMessage.SendMessage,newMessage)

	sliceAllMessage.Send()
	fmt.Println("Hola mundo")

}