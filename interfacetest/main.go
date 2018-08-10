package main

import (
	"thedefy/LearningNotes/interfacetest/user"
	"fmt"
	real2 "thedefy/LearningNotes/interfacetest/real"
	"time"
)

type Notifier interface {
	Notify()
}

type Retriever interface {
	Get(url string) string
}

func main() {

	u := user.Notifier{
		Name:  "ChengLei",
		Email: "TheDefy@163.com",
	}
	sendEmail(&u)

	fmt.Println()

	var r Retriever
	r = &real2.Retriever{
		UserAgent: "",
		TimeOut:   time.Minute,
	}

	downloads := doDownloads(r)
	fmt.Println(downloads)
}

func sendEmail(n Notifier) {
	n.Notify()
}

func doDownloads(r Retriever) string {
	return r.Get("https://github.com/TheDefy/golang-learning-notes")
}
