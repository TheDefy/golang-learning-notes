package main

import "thedefy/LearningNotes/interfacetest/user"

type Notifier interface {
	Notify()
}

func main() {

	u := user.Notifier{
		Name:  "ChengLei",
		Email: "TheDefy@163.com",
	}

	sendEmail(&u)
}

func sendEmail(n Notifier) {
	n.Notify()
}
