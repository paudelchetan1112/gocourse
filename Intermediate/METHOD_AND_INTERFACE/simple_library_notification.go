package main

import "fmt"

type notifier interface {
	sendMessage(message string)
}
type EmailNotifier struct {
	Email string
}
type SMSNotifier struct {
	PhoneNumber string
}

func (e EmailNotifier) sendMessage(message string) {
	fmt.Println("Email:", message)


}
func (s SMSNotifier) sendMessage(message string) {
fmt.Println("SMS:", message)
}
type member struct{
	memberName string
	Notification notifier
	bookName string
}
func (m member) notifyDueBook(){
	message:=fmt.Sprintf("%s, your book %s is due tomorrow.", m.memberName, m.bookName)
	m.Notification.sendMessage(message)
}



func main() {
	//member using email

	member1:=member{memberName: "chetan", Notification: EmailNotifier{Email: "paudelchetan1112@gmail.com"}, bookName: "Go programming"}
	member1.notifyDueBook()
	member2:=member{memberName: "dinesh", Notification: SMSNotifier{PhoneNumber: "98344353"}, bookName: "physic"}
	member2.notifyDueBook()

}