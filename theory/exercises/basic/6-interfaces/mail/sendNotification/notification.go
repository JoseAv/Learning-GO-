package sendnotification

import "fmt"

type Notification interface {
	email(m Message)
	sms(m Message)
	push(m Message)
	Send()
}

type Message struct {
	Title       string
	Description string
}

type AllMessage struct {
	SendMessage []Message
}

func (a *AllMessage) email(m Message) {
	fmt.Println("Email dice :",m.Title)
}

func (a *AllMessage) sms(m Message) {
	fmt.Println("SMS dice :",m.Title)

}

func (a *AllMessage) push(m Message) {
	fmt.Println("Push dice :",m.Title)
}

func (a *AllMessage) Send(){
	for _, val :=range a.SendMessage{
		a.email(val)
		a.sms(val)
		a.push(val)
	}
}