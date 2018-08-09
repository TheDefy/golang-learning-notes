package user

import "fmt"

type Notifier struct {
	Name string
	Email string
}

func (user *Notifier) Notify() {
	fmt.Printf("发送邮件给%s<%s>\n", user.Name, user.Email)
}





