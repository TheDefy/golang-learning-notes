package main

import (
	"thedefy/LearningNotes/interfacetest/user"
	"fmt"
	"thedefy/LearningNotes/interfacetest/mock"
)

type Notifier interface {
	Notify()
}

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string) string
}

type session interface {
	Retriever
	Poster
}

func main() {

	u := user.Notifier{
		Name:  "ChengLei",
		Email: "TheDefy@163.com",
	}
	sendEmail(&u)

	fmt.Println()

	// TODO close
	//var r Retriever
	//r = &real2.Retriever{
	//	UserAgent: "",
	//	TimeOut:   time.Minute,
	//}
	//downloads := doDownloads(r)
	//fmt.Println(downloads)

	fmt.Println()

	//retriever := &mock.Retriever{
	//	Contents: "TheDefy",
	//}
	//
	//doSession(retriever)
	doSession(&mock.Retriever{
		Contents: "TheDefy",
	})
}

func sendEmail(n Notifier) {
	n.Notify()
}

func doDownloads(r Retriever) string {
	return r.Get("https://github.com/TheDefy/golang-learning-notes")
}

func doPoster(r Poster) string {
	return r.Post("https://www.baidu.com")
}

func doSession(s session) string {
	s.Get("wwww.baidu.com")
	return s.Post("wwww.baidu.com")
}
