package hello

import "fmt"

type notifier interface {
	notify2()
}

type user2 struct {
	name  string
	email string
}

func (user *user2) notify2() {
	fmt.Printf("发送邮件给%s<%s>\n", user.name, user.email)
}

func main() {

	u := user2{"chengLei", "thedefy@163.com"}

	//sendEmail(u)
	(&u).notify2()
}
func sendEmail(u notifier) {
	u.notify2()
}
